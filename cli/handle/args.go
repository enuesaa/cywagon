package handle

import "fmt"

func (h *Handler) ValidateArgs(args []string) (error) {
	if len(args) == 0 {
		return fmt.Errorf("required arguments missing: path")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments found")
	}
	return nil
}
