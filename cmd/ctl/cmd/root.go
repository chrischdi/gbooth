package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/chrischdi/gbooth/cmd/ctl/cmd/core"
	"github.com/chrischdi/gbooth/cmd/ctl/cmd/dslr"
	"github.com/chrischdi/gbooth/cmd/ctl/cmd/ui"
)

var rootCmd = &cobra.Command{
	Use: "gboothctl",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(dslr.DSLRCmd)
	rootCmd.AddCommand(ui.UICmd)
	rootCmd.AddCommand(core.CoreCmd)
}
