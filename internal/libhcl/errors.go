package libhcl

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
)

func NewErrParseFailed(err *hcl.Diagnostic) *ErrParseFailed {
	return &ErrParseFailed{
		filename: err.Subject.Filename,
		line:     err.Subject.Start.Line,
		message:  err.Detail,
	}
}

type ErrParseFailed struct {
	filename string
	line     int
	message  string
}

func (e *ErrParseFailed) Error() string {
	return fmt.Sprintf("%s -- %s:%d", e.message, e.filename, e.line)
}
