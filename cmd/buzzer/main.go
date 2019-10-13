package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chrischdi/gbooth/pkg/buzzer"
	"github.com/chrischdi/gbooth/pkg/proto"
	"google.golang.org/grpc"

	"github.com/chrischdi/gbooth/pkg/constants"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().String("core", constants.CoreListen, "address to reach the core grpc server")
	rootCmd.PersistentFlags().String("driver", "gpio", "driver to use for buzzer press detection")
}

var rootCmd = &cobra.Command{
	Use: "dslr",
	RunE: func(cmd *cobra.Command, args []string) error {

		buz, err := buzzer.NewBuzzer(cmd.Flag("driver").Value.String(), buzzer.NewDefaultoptions())
		if err != nil {
			return err
		}

		for {
			if err := loop(buz, cmd.Flag("core").Value.String()); err != nil {
				log.Printf("error detected: %v", err)
			}
		}

		return fmt.Errorf("this should have never been reached")
	},
}

func loop(buz buzzer.Buzzer, core string) error {
	for {
		if buz.Pressed() {
			connCore, err := grpc.Dial(core, grpc.WithInsecure(), grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect to dslr service: %v", err)
			}
			defer connCore.Close()
			coreClient := proto.NewCoreClient(connCore)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()
			if _, err := coreClient.Trigger(ctx, &proto.TriggerRequest{}); err != nil {
				return fmt.Errorf("error on Trigger: %v", err)
			}
			continue
		}
		err := buz.Wait()
		if err != nil {
			return fmt.Errorf("buzzer wait error: %v", err)
		}
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
