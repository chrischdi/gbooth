package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/chrischdi/gbooth/pkg/constants"
	"github.com/chrischdi/gbooth/pkg/ui"
	"github.com/spf13/cobra"

	"github.com/chrischdi/gbooth/pkg/proto"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.PersistentFlags().String("listen", constants.UIListen, "address to listen on for grpc")
	rootCmd.PersistentFlags().String("driver", "gphoto2", "driver to use for dslr connection")
}

var rootCmd = &cobra.Command{
	Use: "dslr",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("starting grpc server on %s", cmd.Flag("listen").Value.String())
		lis, err := net.Listen("tcp", cmd.Flag("listen").Value.String())
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		drv, err := ui.NewUI(cmd.Flag("driver").Value.String())
		if err != nil {
			return err
		}

		if err := drv.Init(); err != nil {
			return err
		}

		s := grpc.NewServer()
		proto.RegisterUIServer(s, drv)
		if err := s.Serve(lis); err != nil {
			return fmt.Errorf("error serving: %v", err)
		}

		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func (s *server) Background(ctx context.Context, in *proto.Image) (*proto.UIResponse, error) {
// 	return nil, fmt.Errorf("background unimplemented")
// }

// func (s *server) Timer(ctx context.Context, in *proto.TextRequest) (*proto.UIResponse, error) {
// 	return nil, fmt.Errorf("timer unimplemented")
// }

// func (s *server) Error(ctx context.Context, in *proto.TextRequest) (*proto.UIResponse, error) {
// 	return nil, fmt.Errorf("error unimplemented")
// }

// func main() {
// 	flag.Parse()

// 	log.Printf("starting grpc server on %s", *listen)
// 	lis, err := net.Listen("tcp", *listen)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	drv, err := ui.NewUI(*driver)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	if err := drv.Init(); err != nil {
// 		if err := drv.Free(); err != nil {
// 			log.Printf("warning: %v", err)
// 		}
// 		log.Fatalf(err.Error())
// 	}

// 	s := grpc.NewServer()
// 	proto.RegisterUIServer(s, &server{})
// 	if err := s.Serve(lis); err != nil {
// 		if err := drv.Free(); err != nil {
// 			log.Printf("warning: %v", err)
// 		}
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
