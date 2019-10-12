package core

import (
	"fmt"

	"github.com/chrischdi/gbooth/pkg/constants"
	"github.com/spf13/cobra"
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
		return fmt.Errorf("unimplemented: connecting to: %s", cmd.Flag("grpc").Value)
	},
}
