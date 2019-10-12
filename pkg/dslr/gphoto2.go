package dslr

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"path"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chrischdi/gbooth/pkg/proto"
)

type GPhoto2 struct {
	mutex           sync.Mutex
	cmd             cmdExecutor
	targetDirectory string
}

func NewGPhoto2() *GPhoto2 {
	return &GPhoto2{}
}

func (s *GPhoto2) Capture(ctx context.Context, in *proto.Request) (*proto.ImageResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	stdout, stderr, err := s.cmd.execute([]string{"gphoto2", "--capture-image-and-download", "--filename", path.Join(s.targetDirectory, "%Y-%m-%d %H-%M-%S.%C")})
	if err != nil {
		return nil, status.Error(codes.Internal, stderr)
	}

	lines := strings.Split(stdout, "\n")
	if len(lines) == 4 {
		fmt.Printf("lines[1]: >%s<\n", lines[1])
		fmt.Printf("trimmed: >%s<\n", strings.TrimPrefix(lines[1], "Saving file as "))
		return &proto.ImageResponse{Name: strings.TrimPrefix(lines[1], "Saving file as ")}, nil
	}

	return nil, status.Error(codes.Unknown, stderr)
}

func (s *GPhoto2) GetDate(ctx context.Context, in *proto.Request) (*proto.DateResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	stdout, stderr, err := s.cmd.execute([]string{"gphoto2", "--get-config", "/main/settings/datetime"})
	if err != nil {
		return nil, status.Error(codes.Internal, strings.Split(stderr, "\n")[0])
	}

	lines := strings.Split(stdout, "\n")
	if len(lines) != 9 {
		return nil, status.Errorf(codes.OutOfRange, "expected 9 lines on stdout, have %s", len(lines))
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	timeStamp, err := time.Parse("Mon 2 Jan 2006 15:04:05 AM MST", strings.TrimPrefix(lines[4], "Printable: "))

	return &proto.DateResponse{
		Date: timeStamp.String(),
	}, nil
}

type cmdExecutor struct{}

func (c *cmdExecutor) execute(cmdline []string) (string, string, error) {
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
