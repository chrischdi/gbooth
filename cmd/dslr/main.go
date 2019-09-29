package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/chrischdi/gbooth/pkg/dslr"
	"github.com/chrischdi/gbooth/pkg/proto"
	"google.golang.org/grpc"
)

// server is used to implement proto.CameraServer.
type server struct {
	dslr dslr.DSLR
}

var (
	port   = flag.String("listen", "127.0.0.1:50051", "address to listen on for grpc")
	driver = flag.String("driver", "canon", "driver to use for dslr connection")
)

// Capture implements helloworld.GreeterServer
func (s *server) Capture(ctx context.Context, in *proto.ImageRequest) (*proto.ImageResponse, error) {
	img, _, err := s.dslr.Capture()
	if err != nil {
		return nil, fmt.Errorf("error during dslr.Capture: %w", err)
	}
	return &proto.ImageResponse{Image: img, Format: "jpeg"}, nil
}

func main() {
	log.Printf("starting grpc server on %s", *port)
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	drv, err := dslr.NewDSLR(*driver)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := drv.Init(); err != nil {
		if err := drv.Free(); err != nil {
			log.Printf("warning: %v", err)
		}
		log.Fatalf(err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterDSLRServer(s, &server{drv})
	if err := s.Serve(lis); err != nil {
		if err := drv.Free(); err != nil {
			log.Printf("warning: %v", err)
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
