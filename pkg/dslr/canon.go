package dslr

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/chrischdi/gphoto2go"
)

type Canon struct {
	cam   gphoto2go.Camera
	mutex sync.Mutex
}

func NewCanon() *Canon {
	c := Canon{
		cam: gphoto2go.Camera{},
	}
	return &c
}

func (c *Canon) Init() error {
	return nil
}

func (c *Canon) Free() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if err := c.cam.Exit(); err != nil {
		return fmt.Errorf("error executing gphoto2go.Exit(): %v", err)
	}
	return nil
}

func (c *Canon) Capture() (data []byte, filename string, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fp, err := c.cam.TriggerCaptureToFile()
	if err != nil {
		return nil, "", fmt.Errorf("TriggerCaptureToFile: %v", err)
	}
	cameraFileReader := c.cam.FileReader(fp.Folder, fp.Name)
	defer cameraFileReader.Close()

	buf, err := ioutil.ReadAll(cameraFileReader)
	if err != nil {
		return nil, "", fmt.Errorf("Error on ioutil ReadAll")
	}

	return buf, fp.Name, nil
}
