package dslr

import (
	"context"
	"fmt"

	"github.com/chrischdi/gbooth/pkg/proto"
)

type DSLR interface {
	Capture(ctx context.Context, in *proto.Request) (*proto.ImageResponse, error)
	GetDate(ctx context.Context, in *proto.Request) (*proto.DateResponse, error)
}

func NewDSLR(driver string) (DSLR, error) {
	switch driver {
	// case "canon":
	// 	return NewCanon(), nil
	case "gphoto2":
		return NewGPhoto2(), nil
	case "dummy":
		return nil, fmt.Errorf("driver %q unimplemented", driver)
	}
	return nil, fmt.Errorf("driver %q not found", driver)
}
