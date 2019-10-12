package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/chrischdi/gbooth/pkg/constants"
	"github.com/chrischdi/gbooth/pkg/dslr"
	"github.com/chrischdi/gbooth/pkg/proto"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().String("listen", constants.DSLRListen, "address to listen on for grpc")
	rootCmd.PersistentFlags().String("driver", "gphoto2", "driver to use for dslr connection")
}

var rootCmd = &cobra.Command{
	Use: "gboothctl",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("starting grpc server on %s", cmd.Flag("listen").Value.String())
		lis, err := net.Listen("tcp", cmd.Flag("listen").Value.String())
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		drv, err := dslr.NewDSLR(cmd.Flag("driver").Value.String())
		if err != nil {
			return err
		}

		s := grpc.NewServer()
		proto.RegisterDSLRServer(s, drv)
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
