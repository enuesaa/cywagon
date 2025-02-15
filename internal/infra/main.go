package infra

var I = Container{
	Fs:  &Fs{},
	Ps:  &Ps{},
	Log: &Log{},
	Cmd: &Cmd{},
}

type Container struct {
	Fs  FsInterface
	Ps  PsInterface
	Log LogInterface
	Cmd CmdInterface
}
