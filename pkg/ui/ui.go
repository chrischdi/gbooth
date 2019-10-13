package ui

import (
	"context"
	"fmt"

	"github.com/chrischdi/gbooth/pkg/proto"
	"github.com/chrischdi/gbooth/pkg/ui/cli"
	"github.com/chrischdi/gbooth/pkg/ui/gtk"
)

type UI interface {
	Init() error
	Free() error
	Background(ctx context.Context, in *proto.Image) (*proto.UIResponse, error)
	Timer(ctx context.Context, in *proto.TextRequest) (*proto.UIResponse, error)
	Error(ctx context.Context, in *proto.TextRequest) (*proto.UIResponse, error)
}

func NewUI(driver string) (UI, error) {
	switch driver {
	case "gtk":
		return gtk.NewGTK()
	case "cli":
		return cli.NewCLI()
	}
	return nil, fmt.Errorf("driver %q not found", driver)
}
