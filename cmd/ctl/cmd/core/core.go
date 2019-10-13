package core

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

var CoreCmd = &cobra.Command{
	Use: "core",
}

func init() {
	CoreCmd.PersistentFlags().String("grpc", constants.CoreListen, "addr to reach core grpc service")
	CoreCmd.AddCommand(triggerCmd)
}

var triggerCmd = &cobra.Command{
	Use: "trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(cmd.Flag("grpc").Value.String(), grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("did not connect: %v", err)
		}
		defer conn.Close()

		c := proto.NewCoreClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		if _, err := c.Trigger(ctx, &proto.TriggerRequest{}); err != nil {
			return fmt.Errorf("could not trigger: %v", err)
		}

		log.Printf("done triggering")

		return nil
	},
}
