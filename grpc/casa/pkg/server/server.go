package server

import (
	"casa-api/grpc/casa/pkg/communication"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const tokenHeader = "x-communication-token"

type ClientMessage struct {
	UserID  string
	Message *communication.StreamResponse
}

type server struct {
	Host, Password string

	StreamServers map[string]communication.Casa_StreamServer
	userIDs       map[string]string
	Messages      chan ClientMessage

	userIDsMtx, streamServersMtx sync.RWMutex
	tokenRemove                  string
}

func Server(host, pass string) *server {
	return &server{
		Host:     host,
		Password: pass,

		userIDs: make(map[string]string),

		StreamServers: make(map[string]communication.Casa_StreamServer),

		Messages: make(chan ClientMessage),
	}
}

func (s *server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	logrus.Info("starting on %s with password %q",
		s.Host, s.Password)

	srv := grpc.NewServer()
	communication.RegisterCasaServer(srv, s)

	l, err := net.Listen("tcp", s.Host)
	if err != nil {
		return errors.WithMessage(err,
			"server unable to bind on provided host")
	}

	go func() {
		for {
			time.Sleep(time.Second * 1)
			logrus.Info("Send to client")
			alarms := &communication.StreamResponse_Alarms{Id: "test", Identifier: "alarms", Method: "test"}

			res := &communication.StreamResponse{
				Event:     &communication.StreamResponse_ClientAlarms{ClientAlarms: alarms},
				Timestamp: ptypes.TimestampNow(),
			}

			m := ClientMessage{
				UserID:  "hubba",
				Message: res,
			}
			s.Messages <- m
		}
	}()

	go func() {
		srv.Serve(l)
		cancel()
	}()

	go func() {
		s.listenToMessage(ctx)
	}()

	<-ctx.Done()

	// s.Broadcast <- communication.StreamResponse{
	// 	Timestamp: ptypes.TimestampNow(),
	// 	Event: &communication.StreamResponse_ServerShutdown{
	// 		&communication.StreamResponse_Shutdown{}}}

	// close(s.Broadcast)
	// ServerLogf(time.Now(), "shutting down")

	srv.GracefulStop()
	return nil
}

func (s *server) listenToMessage(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logrus.Info("Finished listen to messages")
			return
		case m := <-s.Messages:
			logrus.Info("Sending message to client")
			streamServer, ok := s.getStreamServer(m.UserID)
			if ok {
				err := streamServer.Send(m.Message)
				if err != nil {
					logrus.Warnf("Failed to send message to client: %s, %s", m.UserID, err)
				} else {
					logrus.Infof("Send message to %s", m.UserID)
				}
			} else {
				logrus.Infof("Could not find server for:%s", m.UserID)
			}
		}
	}
}

func (s *server) Login(ctx context.Context, req *communication.LoginRequest) (*communication.LoginResponse, error) {
	switch {
	case req.Password != s.Password:
		return nil, status.Error(codes.Unauthenticated, "password is incorrect")
	case req.Name == "":
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	tkn := s.genToken()
	s.setUserID(tkn, req.Name)
	s.tokenRemove = tkn

	logrus.Infof("%s (%s) has logged in", tkn, req.Name)

	// s.Broadcast <- communication.StreamResponse{
	// 	Timestamp: ptypes.TimestampNow(),
	// 	Event: &communication.StreamResponse_Login{
	// 		Name: req.Name,
	// 	},
	// }

	return &communication.LoginResponse{Token: tkn}, nil
}

func (s *server) Logout(ctx context.Context, req *communication.LogoutRequest) (*communication.LogoutResponse, error) {
	ok, userID := s.delUserID(req.Token)
	if !ok {
		return nil, status.Error(codes.NotFound, "token not found")
	}

	logrus.Infof("%s (%s) has logged out", req.Token, userID)

	return new(communication.LogoutResponse), nil
}

func (s *server) sendAlarms(ctx context.Context, userID string, message *communication.StreamResponse_Alarms) (bool, error) {
	res := &communication.StreamResponse{
		Event:     &communication.StreamResponse_ClientAlarms{ClientAlarms: message},
		Timestamp: ptypes.TimestampNow(),
	}

	streamServer, ok := s.getStreamServer(userID)
	if ok {
		return true, streamServer.Send(res)
	}
	return false, nil
}

func (s *server) Stream(srv communication.Casa_StreamServer) error {
	tkn, ok := s.extractToken(srv.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "missing token header")
	}

	userID, ok := s.getUserID(tkn)
	if !ok {
		return status.Error(codes.Unauthenticated, "invalid token")
	}

	s.setStreamServer(userID, srv)

	for {
		req, err := srv.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			logrus.Warnf("GRPC stream stopped: %s", err)
			s.delStreamServer(userID)
			s.delUserID(userID)
			return err
		}

		id := req.GetId()
		logrus.Infof("Received message with id: %s", id)

		if id == "/alarms" {
			alarms := &communication.StreamResponse_Alarms{Id: id, Identifier: "alarms", Method: "test"}
			s.sendAlarms(srv.Context(), tkn, alarms)
		} else {
			logrus.Infof("unknown stream id: %s", id)
		}
		// TODO SEND Message to client
		// s.Broadcast <- communication.StreamResponse{
		// 	Timestamp: ptypes.TimestampNow(),
		// 	Event: &communication.StreamResponse_ ClientMessage{&communication.StreamResponse_ClientMessage{
		// 		Name:    name,
		// 		Message: req.Message,
		// 	}},
		// }
	}

	<-srv.Context().Done()
	return srv.Context().Err()
}

func (s *server) genToken() string {
	tkn := make([]byte, 4)
	rand.Read(tkn)
	return fmt.Sprintf("%x", tkn)
}

func (s *server) getUserID(token string) (UserID string, ok bool) {
	s.userIDsMtx.RLock()
	UserID, ok = s.userIDs[token]
	s.userIDsMtx.RUnlock()
	return
}

func (s *server) setUserID(token string, userID string) {
	s.userIDsMtx.Lock()
	s.userIDs[token] = userID
	s.userIDsMtx.Unlock()
}

func (s *server) delUserID(token string) (ok bool, userID string) {
	userID, ok = s.getUserID(token)

	if ok {
		s.userIDsMtx.Lock()
		delete(s.userIDs, token)
		s.userIDsMtx.Unlock()
	}

	return
}

func (s *server) getStreamServer(userID string) (streamServer communication.Casa_StreamServer, ok bool) {
	s.streamServersMtx.RLock()
	streamServer, ok = s.StreamServers[userID]
	s.streamServersMtx.RUnlock()
	return
}

func (s *server) setStreamServer(userID string, streamServer communication.Casa_StreamServer) {
	s.streamServersMtx.Lock()
	s.StreamServers[userID] = streamServer
	s.streamServersMtx.Unlock()
}

func (s *server) delStreamServer(userID string) (ok bool) {
	_, ok = s.getStreamServer(userID)

	if ok {
		s.streamServersMtx.Lock()
		delete(s.StreamServers, userID)
		s.streamServersMtx.Unlock()
	}

	return
}

func (s *server) extractToken(ctx context.Context) (tkn string, ok bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md[tokenHeader]) == 0 {
		return "", false
	}

	return md[tokenHeader][0], true
}
