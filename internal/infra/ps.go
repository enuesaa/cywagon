package infra

import "os"

type PsInterface interface {
	Exit(code int)
}
type Ps struct{}

func (i *Ps) Exit(code int) {
	os.Exit(code)
}
