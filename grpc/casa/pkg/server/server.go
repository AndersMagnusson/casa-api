package server

import (
	"casa-api/grpc/casa/pkg/communication"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"

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

type server struct {
	Host, Password string

	Broadcast chan communication.StreamResponse

	ClientNames   map[string]string
	ClientStreams map[string]chan communication.StreamResponse

	namesMtx, streamsMtx sync.RWMutex
	tokenRemove          string
}

func Server(host, pass string) *server {
	return &server{
		Host:     host,
		Password: pass,

		Broadcast: make(chan communication.StreamResponse, 1000),

		ClientNames:   make(map[string]string),
		ClientStreams: make(map[string]chan communication.StreamResponse),
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

	go s.broadcast(ctx)

	go func() {
		srv.Serve(l)
		cancel()
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

func (s *server) Login(ctx context.Context, req *communication.LoginRequest) (*communication.LoginResponse, error) {
	switch {
	case req.Password != s.Password:
		return nil, status.Error(codes.Unauthenticated, "password is incorrect")
	case req.Name == "":
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	tkn := s.genToken()
	s.setName(tkn, req.Name)
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
	name, ok := s.delName(req.Token)
	if !ok {
		return nil, status.Error(codes.NotFound, "token not found")
	}

	logrus.Infof("%s (%s) has logged out", req.Token, name)

	// s.Broadcast <- communication.StreamResponse{
	// 	Timestamp: ptypes.TimestampNow(),
	// 	Event: &communication.StreamResponse_Logout{
	// 		Name: name,
	// 	},
	// }

	return new(communication.LogoutResponse), nil
}

func (s *server) sendAlarms(ctx context.Context, token string, message *communication.StreamResponse_Alarms) {
	res := communication.StreamResponse{
		Event:     &communication.StreamResponse_ClientAlarms{ClientAlarms: message},
		Timestamp: ptypes.TimestampNow(),
	}

	s.sendMessage(ctx, token, res)
}

func (s *server) Stream(srv communication.Casa_StreamServer) error {
	tkn, ok := s.extractToken(srv.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "missing token header")
	}

	_, ok = s.getName(tkn)
	if !ok {
		return status.Error(codes.Unauthenticated, "invalid token")
	}

	go s.sendBroadcasts(srv, tkn)

	for {
		req, err := srv.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		id := req.GetId()

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

func (s *server) sendBroadcasts(srv communication.Casa_StreamServer, tkn string) {
	stream := s.openStream(tkn)
	s.sendAlarms(context.Background(), "tokenRemove", &communication.StreamResponse_Alarms{Id: "hej", Identifier: "alarms", Method: "get"})
	defer s.closeStream(tkn)

	for {
		select {
		case <-srv.Context().Done():
			return
		case res := <-stream:
			if s, ok := status.FromError(srv.Send(&res)); ok {
				switch s.Code() {
				case codes.OK:
					// noop
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					logrus.Infof("client (%s) terminated connection", tkn)
					return
				default:
					logrus.Infof("failed to send to client (%s): %v", tkn, s.Err())
					return
				}
			}
		}
	}
}

func (s *server) broadcast(ctx context.Context) {
	for res := range s.Broadcast {
		s.streamsMtx.RLock()
		for _, stream := range s.ClientStreams {
			select {
			case stream <- res:
				// noop
			default:
				logrus.Infof("client stream full, dropping message")
			}
		}
		s.streamsMtx.RUnlock()
	}
}

func (s *server) sendMessage(ctx context.Context, clientID string, res communication.StreamResponse) {
	s.streamsMtx.RLock()
	for id, stream := range s.ClientStreams {
		if id == clientID {
			select {
			case stream <- res:
				// noop
			default:
				logrus.Infof("client stream full, dropping message")
			}
			break
		}
	}
	s.streamsMtx.RUnlock()
}

func (s *server) openStream(tkn string) (stream chan communication.StreamResponse) {
	stream = make(chan communication.StreamResponse, 100)

	s.streamsMtx.Lock()
	s.ClientStreams[tkn] = stream
	s.streamsMtx.Unlock()

	logrus.Infof("opened stream for client %s", tkn)

	return
}

func (s *server) closeStream(tkn string) {
	s.streamsMtx.Lock()

	if stream, ok := s.ClientStreams[tkn]; ok {
		delete(s.ClientStreams, tkn)
		close(stream)
	}

	logrus.Infof("closed stream for client %s", tkn)

	s.streamsMtx.Unlock()
}

func (s *server) genToken() string {
	tkn := make([]byte, 4)
	rand.Read(tkn)
	return fmt.Sprintf("%x", tkn)
}

func (s *server) getName(tkn string) (name string, ok bool) {
	s.namesMtx.RLock()
	name, ok = s.ClientNames[tkn]
	s.namesMtx.RUnlock()
	return
}

func (s *server) setName(tkn string, name string) {
	s.namesMtx.Lock()
	s.ClientNames[tkn] = name
	s.namesMtx.Unlock()
}

func (s *server) delName(tkn string) (name string, ok bool) {
	name, ok = s.getName(tkn)

	if ok {
		s.namesMtx.Lock()
		delete(s.ClientNames, tkn)
		s.namesMtx.Unlock()
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
