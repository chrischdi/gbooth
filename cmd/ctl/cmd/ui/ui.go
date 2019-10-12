package ui

import (
	"fmt"

	"github.com/chrischdi/gbooth/pkg/constants"
	"github.com/spf13/cobra"
)

var UICmd = &cobra.Command{
	Use: "ui",
}

func init() {
	UICmd.PersistentFlags().String("grpc", constants.UIListen, "addr to reach ui grpc service")
	UICmd.AddCommand(backgroundCmd)
	UICmd.AddCommand(timerCmd)
	UICmd.AddCommand(errorCmd)
}

var backgroundCmd = &cobra.Command{
	Use: "background",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("unimplemented: connecting to: %s", cmd.Flag("grpc").Value)
	},
}

var timerCmd = &cobra.Command{
	Use: "timer",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("unimplemented")
	},
}

var errorCmd = &cobra.Command{
	Use: "error",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("unimplemented")
	},
}
