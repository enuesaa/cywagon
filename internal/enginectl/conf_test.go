package enginectl

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
)

func TestValidateConfs(t *testing.T) {
	engine := New()
	engine.Container = infra.NewMock(t)

	table := []struct {
		confs []model.Conf
		err   error
	}{
		{
			confs: []model.Conf{
				{
					Host: "",
				},
			},
			err: service.ErrConfHostRequired,
		},
	}

	for _, tt := range table {
		err := engine.ValidateConfs(tt.confs)
		assert.Equal(t, err, tt.err)
	}
}
