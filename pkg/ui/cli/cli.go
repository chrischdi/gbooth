package cli

import (
	"context"
	"fmt"

	"github.com/chrischdi/gbooth/pkg/proto"
)

const (
	Title = "Photobox"
)

func NewCLI() (*CLI, error) {
	ui := &CLI{}
	return ui, nil
}

type CLI struct{}

func (ui *CLI) Init() error {
	return nil
}

func (ui *CLI) Free() error {
	return nil
}

func (ui *CLI) Background(ctx context.Context, in *proto.Image) (*proto.UIResponse, error) {
	fmt.Printf("publishing background\n")
	return &proto.UIResponse{}, nil
}

func (ui *CLI) Timer(ctx context.Context, in *proto.TextRequest) (*proto.UIResponse, error) {
	fmt.Printf("publishing timer: %q\n", in.GetText())
	return &proto.UIResponse{}, nil
}

func (ui *CLI) Error(ctx context.Context, in *proto.TextRequest) (*proto.UIResponse, error) {
	fmt.Printf("publishing error: %q\n", in.GetText())
	return &proto.UIResponse{}, nil
}
