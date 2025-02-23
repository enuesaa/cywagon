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

type Mock struct {
	Fs  *MockFsInterface
	Ps  *MockPsInterface
	Log *MockLogInterface
	Cmd *MockCmdInterface
}

func (m *Mock) Container() Container {
	return Container{
		Fs:  m.Fs,
		Ps:  m.Ps,
		Log: m.Log,
		Cmd: m.Cmd,
	}
}

func (m *Mock) With(prepare func(*Mock)) Container {
	prepare(m)

	return m.Container()
}

func NewMock(t *testing.T) *Mock {
	ctrl := gomock.NewController(t)

	return &Mock{
		Fs:  NewMockFsInterface(ctrl),
		Ps:  NewMockPsInterface(ctrl),
		Log: NewMockLogInterface(ctrl),
		Cmd: NewMockCmdInterface(ctrl),
	}
}
