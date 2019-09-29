package dslr

import "fmt"

type DSLR interface {
	Init() error
	Free() error
	Capture() (data []byte, filename string, err error)
}

func NewDSLR(driver string) (DSLR, error) {
	switch driver {
	case "canon":
		return NewCanon(), nil
	case "dummy":
		return nil, fmt.Errorf("driver %q unimplemented", driver)
	}
	return nil, fmt.Errorf("driver %q not found", driver)
}
