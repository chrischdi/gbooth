package ui

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/chrischdi/gbooth/pkg/constants"
	"github.com/chrischdi/gbooth/pkg/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var UICmd = &cobra.Command{
	Use: "ui",
}

func init() {
	backgroundCmd.PersistentFlags().String("image", "", "path to an jpg file (max size <4MB)")
	UICmd.PersistentFlags().String("grpc", constants.UIListen, "addr to reach ui grpc service")
	UICmd.AddCommand(backgroundCmd)
	UICmd.AddCommand(timerCmd)
	UICmd.AddCommand(errorCmd)
}

var backgroundCmd = &cobra.Command{
	Use: "background",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := cmd.Flag("image").Value.String()
		if path == "" {
			return fmt.Errorf("the `image` parameter is mandatory")
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		conn, err := grpc.Dial(cmd.Flag("grpc").Value.String(), grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("did not connect: %v", err)
		}
		defer conn.Close()

		c := proto.NewUIClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
		defer cancel()

		if _, err := c.Background(ctx, &proto.Image{Image: b}); err != nil {
			return err
		}
		return nil
	},
}

var timerCmd = &cobra.Command{
	Use: "timer",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(cmd.Flag("grpc").Value.String(), grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("did not connect: %v", err)
		}
		defer conn.Close()

		c := proto.NewUIClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
		defer cancel()

		if _, err := c.Timer(ctx, &proto.TextRequest{Text: strings.Join(args, " ")}); err != nil {
			return err
		}
		return nil
	},
}

var errorCmd = &cobra.Command{
	Use: "error",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(cmd.Flag("grpc").Value.String(), grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("did not connect: %v", err)
		}
		defer conn.Close()

		c := proto.NewUIClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
		defer cancel()

		if _, err := c.Error(ctx, &proto.TextRequest{Text: strings.Join(args, " ")}); err != nil {
			return err
		}
		return nil
	},
}
