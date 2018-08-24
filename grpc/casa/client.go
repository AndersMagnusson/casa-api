package casa

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"casa-api/grpc/casa/protos"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type client struct {
	communication.CasaClient
	Host, Password, Name, Token string
	Shutdown                    bool
}

func Client(host, pass, name string) *client {
	return &client{
		Host:     host,
		Password: pass,
		Name:     name,
	}
}

func (c *client) Run(ctx context.Context) error {
	connCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.DialContext(connCtx, c.Host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.WithMessage(err, "unable to connect")
	}
	defer conn.Close()

	c.CasaClient = communication.NewCasaClient(conn)

	if c.Token, err = c.login(ctx); err != nil {
		return errors.WithMessage(err, "failed to login")
	}
	entry := logrus.WithField("grpc", "run")

	entry.Info("Logged in successfully")
	entry.Info("logged in successfully")

	err = c.stream(ctx)

	entry.Info("Logging out")
	if err := c.logout(ctx); err != nil {
		entry.Infof("failed to log out: %v", err)
	}

	return errors.WithMessage(err, "stream error")
}

func (c *client) stream(ctx context.Context) error {
	entry := logrus.WithField("grpc", "stream")
	md := metadata.New(map[string]string{tokenHeader: c.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := c.CasaClient.Stream(ctx)
	if err != nil {
		return err
	}
	defer client.CloseSend()

	entry.Info("connected to stream")

	go c.send(client)
	return c.receive(client)
}

func (c *client) receive(sc communication.Casa_StreamClient) error {
	entry := logrus.WithField("grpc", "receive")
	for {
		res, err := sc.Recv()

		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			entry.Info("stream canceled (usually indicates shutdown)")
			return nil
		} else if err == io.EOF {
			entry.Info("stream closed by server")
			return nil
		} else if err != nil {
			return err
		}

		// ts := tsToTime(res.Timestamp)

		switch evt := res.Event.(type) {
		case *communication.StreamResponse_ClientLogin:
			// TODO Handle login
			// ServerLogf(ts, "%s has logged in", evt.ClientLogin.Name)
		case *communication.StreamResponse_ClientAlarms:
			// ServerLogf(ts, "%s has logged out", evt.ClientLogout.Name)
		case *communication.StreamResponse_ClientAlerts:

			// MessageLog(ts, evt.ClientMessage.Name, evt.ClientMessage.Message)

		default:
			entry.Infof("unexpected event from the server: %T", evt)
			return nil
		}
	}
}

func (c *client) send(client communication.Casa_StreamClient) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	entry := logrus.WithField("grpc", "send")
	for {
		select {
		case <-client.Context().Done():
			entry.Debug("client send loop disconnected")
		default:
			if sc.Scan() {
				if err := client.Send(&communication.StreamRequest{Message: sc.Text()}); err != nil {
					entry.Debug("failed to send message: %v", err)
					return
				}
			} else {
				entry.Infof("input scanner failure: %v", sc.Err())
				return
			}
		}
	}
}

func (c *client) login(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := c.CasaClient.Login(ctx, &communication.LoginRequest{
		Name:     c.Name,
		Password: c.Password,
	})

	if err != nil {
		return "", err
	}

	return res.Token, nil
}

func (c *client) logout(ctx context.Context) error {
	if c.Shutdown {
		logrus.Debug("unable to logout (server sent shutdown signal)")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.CasaClient.Logout(ctx, &communication.LogoutRequest{Token: c.Token})
	if s, ok := status.FromError(err); ok && s.Code() == codes.Unavailable {
		logrus.Debug("unable to logout (connection already closed)")
		return nil
	}

	return err
}
