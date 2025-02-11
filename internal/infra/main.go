package infra

func New() Container {
	return Container{
		Fs:  &Fs{},
		Log: &Log{},
		Cmd: &Cmd{},
	}
}

type Container struct {
	Fs  FsInterface
	Log LogInterface
	Cmd CmdInterface
}
