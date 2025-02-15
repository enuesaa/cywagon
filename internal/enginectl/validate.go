package enginectl

import (
	"fmt"

	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Validate(confs []model.Conf) error {
	for _, conf := range confs {
		if conf.Host == "" {
			return fmt.Errorf("host is required")
		}
	}
	return nil
}
