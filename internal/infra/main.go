package infra

import "fmt"


var LogI LogInterface = &Log{}

func New() Container {
	return Container{
		Fs:  &Fs{},
		Ps:  &Ps{},
		Log: &Log{},
		Cmd: &Cmd{},
	}
}

type Container struct {
	Fs  FsInterface
	Ps  PsInterface
	Log LogInterface
	Cmd CmdInterface
}

func SetupMock() {
	fmt.Println("setup mock")
	LogI = &Log{}
}
