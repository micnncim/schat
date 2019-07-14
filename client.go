package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "github.com/micnncim/schat/proto"
)

type Client struct {
	cli  pb.ChatClient
	conn *grpc.ClientConn
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := &Client{
		cli:  pb.NewChatClient(conn),
		conn: conn,
	}
	return c, nil
}

func (c *Client) Run(ctx context.Context) error {
	log.Println("client: starting")

	stream, err := c.cli.SendMessage(ctx)
	if err != nil {
		return err
	}
	defer func() {
		stream.CloseSend()
		log.Println("client: disconnected from stream")
	}()
	log.Println("client: connected to stream")

	u, err := user.Current()
	if err != nil {
		return err
	}

	errCh := make(chan error)
	go func() {
		errCh <- c.send(ctx, stream, u.Username)
	}()
	go func() {
		errCh <- c.receive(ctx, stream)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return errors.New("client: canceled")
	}
}

func (c *Client) send(ctx context.Context, stream pb.Chat_SendMessageClient, username string) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			return errors.New("client: canceled")
		default:
			if scanner.Scan() {
				err := stream.Send(&pb.SendMessageRequest{
					Username: username,
					Message:  scanner.Text(),
				})
				if err != nil {
					return err
				}
				continue
			}
			return errors.New("failed to scan stdin")
		}
	}
}

func (c *Client) receive(ctx context.Context, stream pb.Chat_SendMessageClient) error {
	for {
		fmt.Fprintf(os.Stderr, ">>> ")
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		t, err := ptypes.Timestamp(resp.Timestamp)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "%s %s\n", t.In(defaultLocation).Format(dateFormat), resp.Message)
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}
