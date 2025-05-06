package infra

import (
	"testing"

	"go.uber.org/mock/gomock"
)

var Default = New()

type Container struct {
	Fs  FsInterface
	Ps  PsInterface
	Cmd CmdInterface
}

func New() Container {
	return Container{
		Fs:  &Fs{},
		Ps:  &Ps{},
		Cmd: &Cmd{},
	}
}

type Mock struct {
	Fs  *MockFsInterface
	Ps  *MockPsInterface
	Cmd *MockCmdInterface
}

func (m *Mock) Container() Container {
	return Container{
		Fs:  m.Fs,
		Ps:  m.Ps,
		Cmd: m.Cmd,
	}
}

func NewMock(t *testing.T, prepares ...func(*Mock)) Container {
	ctrl := gomock.NewController(t)

	mock := Mock{
		Fs:  NewMockFsInterface(ctrl),
		Ps:  NewMockPsInterface(ctrl),
		Cmd: NewMockCmdInterface(ctrl),
	}
	for _, prepare := range prepares {
		prepare(&mock)
	}

	return mock.Container()
}
