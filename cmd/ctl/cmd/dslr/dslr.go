package dslr

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chrischdi/gbooth/pkg/constants"

	"github.com/chrischdi/gbooth/pkg/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var DSLRCmd = &cobra.Command{
	Use: "dslr",
}

func init() {
	DSLRCmd.PersistentFlags().String("grpc", constants.DSLRListen, "addr to reach dslr grpc service")
	DSLRCmd.AddCommand(captureCmd)
	DSLRCmd.AddCommand(getDateCmd)
}

var captureCmd = &cobra.Command{
	Use: "capture",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(cmd.Flag("grpc").Value.String(), grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("did not connect: %v", err)
		}
		defer conn.Close()

		c := proto.NewDSLRClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
		defer cancel()

		r, err := c.Capture(ctx, &proto.Request{})
		if err != nil {
			return fmt.Errorf("could not capture image: %v", err)
		}

		log.Printf("done taking picture to %s", r.GetName())

		return nil
	},
}

var getDateCmd = &cobra.Command{
	Use: "getDate",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(cmd.Flag("grpc").Value.String(), grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("did not connect: %v", err)
		}
		defer conn.Close()

		c := proto.NewDSLRClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
		defer cancel()

		r, err := c.GetDate(ctx, &proto.Request{})
		if err != nil {
			return fmt.Errorf("could not get date: %v", err)
		}

		log.Printf("got %s", r.GetDate())

		return nil
	},
}
