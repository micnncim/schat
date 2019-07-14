package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "github.com/micnncim/schat/proto"
)

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	log.Println("server: starting")

	srv := grpc.NewServer()
	pb.RegisterChatServer(srv, s)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.Serve(lis)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return errors.New("server: timeout")
	}
}

func (s *Server) SendMessage(stream pb.Chat_SendMessageServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		now := time.Now().In(defaultLocation)
		ts, err := ptypes.TimestampProto(now)
		if err != nil {
			return err
		}
		println("[debug] ", now.String())
		msg := fmt.Sprintf("%8s | %s", req.Username, req.Message)
		err = stream.Send(&pb.SendMessageResponse{
			Message:   msg,
			Timestamp: ts,
		})
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "%s %s\n", now.Format(dateFormat), msg)
	}
}
