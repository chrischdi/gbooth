package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"gopkg.in/h2non/bimg.v1"

	"github.com/chrischdi/gbooth/pkg/constants"
	"github.com/chrischdi/gbooth/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type server struct {
	b    bool
	m    sync.Mutex
	dslr proto.DSLRClient
	ui   proto.UIClient
}

var (
	listen = flag.String("listen", constants.CoreListen, "address to listen on for grpc")
	dslr   = flag.String("dslr", constants.DSLRListen, "address to reach the dslr grpc server")
	ui     = flag.String("ui", constants.UIListen, "address to reach the ui grpc server")
)

func (s *server) Trigger(ctx context.Context, in *proto.TriggerRequest) (*proto.CoreResponse, error) {
	if err := s.acquireB(true); err != nil {
		return nil, err
	}
	defer func() {
		if err := s.acquireB(false); err != nil {
			log.Printf("error reseting acquireB: %v", err)
		}
	}()

	// countdown
	if err := s.countdown(); err != nil {
		return nil, fmt.Errorf("error on countdown: %v", err)
	}

	// take picture
	ctxDSLR, cancelDSLR := context.WithTimeout(context.Background(), time.Second*7)
	defer cancelDSLR()
	response, err := s.dslr.Capture(ctxDSLR, &proto.Request{})
	if err != nil {
		errMsg := fmt.Sprintf("error on dslr.Capture: %v", err)

		ctxError, cancelError := context.WithTimeout(context.Background(), time.Second*3)
		defer cancelError()
		_, err := s.ui.Error(ctxError, &proto.TextRequest{Text: errMsg})
		if err != nil {
			log.Printf("error showing error: %v", err)
		}

		return nil, status.Error(codes.Internal, errMsg)
	}

	b, err := getResized(response.GetName())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// publish picture
	ctxUI, cancelUI := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelUI()
	if _, err := s.ui.Background(ctxUI, &proto.Image{
		Image: b,
	}); err != nil {
		return nil, fmt.Errorf("error on ui.Background: %v", err)
	}

	return &proto.CoreResponse{}, nil
}

func getResized(path string) ([]byte, error) {
	buffer, err := bimg.Read(path)
	if err != nil {
		return nil, err
	}
	options := bimg.Options{
		Width:  1920,
		Height: 1080,
		Embed:  true,
		Crop:   true,
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		return nil, err
	}
	return newImage, nil
}

func main() {
	flag.Parse()

	log.Printf("starting grpc server on %s", *listen)
	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connDSLR, err := grpc.Dial(*dslr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to dslr service: %v", err)
	}
	defer connDSLR.Close()
	dslrClient := proto.NewDSLRClient(connDSLR)

	connUI, err := grpc.Dial(*ui, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to dslr service: %v", err)
	}
	defer connUI.Close()
	uiClient := proto.NewUIClient(connUI)

	s := grpc.NewServer()
	proto.RegisterCoreServer(s, &server{
		dslr: dslrClient,
		ui:   uiClient,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) countdown() error {
	for i := 3; i > 0; i-- {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		if _, err := s.ui.Timer(ctx, &proto.TextRequest{Text: fmt.Sprintf("%d", i)}); err != nil {
			return fmt.Errorf("error setting countdown: %v", err)
		}
		time.Sleep(time.Second)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	if _, err := s.ui.Timer(ctx, &proto.TextRequest{Text: "Action!"}); err != nil {
		return fmt.Errorf("error setting countdown: %v", err)
	}
	return nil
}

func (s *server) acquireB(b bool) error {
	s.m.Lock()
	defer s.m.Unlock()

	if s.b && b {
		return status.Error(codes.Unavailable, "other request already in progress")
	}

	if !s.b && !b {
		return status.Error(codes.PermissionDenied, "s.b is already set to false")
	}

	s.b = b
	return nil
}
