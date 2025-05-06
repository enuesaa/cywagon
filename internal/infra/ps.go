package infra

import (
	"fmt"
	"os"
)

type PsInterface interface {
	Print(text string)
	Printf(format string, a ...any)
	PrintErr(err error)

	Exit(code int)
}
type Ps struct{}

func (i *Ps) Print(text string) {
	fmt.Println(text)
}

func (i *Ps) Printf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	i.Print(text)
}

func (i *Ps) PrintErr(err error) {
	text := fmt.Sprintf("Error: %s", err.Error())
	i.Print(text)
}

func (i *Ps) Exit(code int) {
	os.Exit(code)
}
