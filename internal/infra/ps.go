package infra

import "os"

// To generate mock file, run following command:
//   mockgen -source=ps.go -destination=ps_mock.go -package=infra

type PsInterface interface {
	Exit(code int)
}
type Ps struct{}

func (i *Ps) Exit(code int) {
	os.Exit(code)
}
