package infra

import (
	"testing"

	"go.uber.org/mock/gomock"
)

var Default = New()

type Container struct {
	Fs  FsInterface
	Ps  PsInterface
	Log LogInterface
	Cmd CmdInterface
}

func New() Container {
	return Container{
		Fs:  &Fs{},
		Ps:  &Ps{},
		Log: &Log{},
		Cmd: &Cmd{},
	}
}

func NewMock(t *testing.T) Container {
	ctrl := gomock.NewController(t)

	return Container{
		Fs:  NewMockFsInterface(ctrl),
		Ps:  NewMockPsInterface(ctrl),
		Log: NewMockLogInterface(ctrl),
		Cmd: NewMockCmdInterface(ctrl),
	}
}
