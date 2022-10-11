package buzzer

import (
	"fmt"
	"time"
)

type Options struct {
	// GPIOPin is the pin used at the gpio interface for detecting button presses
	GPIOPin string
	// File is the file used for the file driver
	File string
	// TrueTimeout is the timeout used between requests for the true buzzer
	TrueTimeout time.Duration
}

// Buzzer represents a trigger which a Photobooth is waiting for to be true
type Buzzer interface {
	// Wait returns as soon as the status of the backend changes
	Wait() error
	// Pressed returns true if the Buzzer is pressed
	Pressed() bool
}

func NewBuzzer(name string, options Options) (Buzzer, error) {
	switch name {
	case "file":
		return &FileBuzzer{File: options.File}, nil
	case "true":
		return &TrueBuzzer{options.TrueTimeout}, nil
	case "gpio":
		b, err := NewGPIOBuzzer(options.GPIOPin)
		if err != nil {
			return nil, fmt.Errorf("failed registering RPI buzzer: %v", err)
		}
		return b, nil
	}
	return nil, fmt.Errorf("buzzer '%s' not found", name)
}

func NewDefaultoptions() Options {
	return Options{
		GPIOPin:     "6",
		File:        "/tmp/buzzer.lock",
		TrueTimeout: time.Second * 3,
	}
}

func HelpString() string {
	return "gpio,file,true"
}
