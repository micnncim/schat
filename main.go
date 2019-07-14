package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	serverMode bool
	port       string
	address    string

	defaultLocation = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func init() {
	flag.BoolVar(&serverMode, "s", false, "run as a server")
	flag.StringVar(&port, "port", ":8080", "port for server")
	flag.StringVar(&address, "address", "localhost:8080", "address of server for client")
	flag.Parse()
}

const dateFormat = "2006/01/02 15:04:05"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if serverMode {
		s := &Server{}
		if err := s.Run(ctx); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	c, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	if err := c.Run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
