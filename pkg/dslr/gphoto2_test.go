package dslr

import (
	"context"
	"reflect"
	"sync"
	"testing"

	"github.com/chrischdi/gbooth/pkg/proto"
)

func TestGPhoto2_Capture(t *testing.T) {
	type fields struct {
		mutex           sync.Mutex
		cmd             gphoto2Cmd
		targetDirectory string
	}
	type args struct {
		ctx context.Context
		in  *proto.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.ImageResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GPhoto2{
				mutex:           tt.fields.mutex,
				cmd:             tt.fields.cmd,
				targetDirectory: tt.fields.targetDirectory,
			}
			got, err := s.Capture(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GPhoto2.Capture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GPhoto2.Capture() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	captureErrStderr = `
*** Error ***              
Canon EOS Capture failed to release: Perhaps no focus?
ERROR: Could not capture image.
ERROR: Could not capture.
`

	captureStdout = `New file is in location /store_00020001/DCIM/100CANON/IMG_2671.JPG on the camera
Saving file as 2019-10-10 20-05-20.JPG
Deleting file /store_00020001/DCIM/100CANON/IMG_2671.JPG on the camera
`

	getDateStdout = `Label: Camera Date and Time
Readonly: 0
Type: DATE
Current: 1570731662
Printable: Thu 10 Oct 2019 08:21:02 PM CEST
Help: Use 'now' as the current time when setting.

END
`

	getDateCameraOffErrStderr = `*** Error: No camera found. ***

`
)
