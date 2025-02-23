package enginectl

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
)

func TestValidateConfs(t *testing.T) {
	engine := New(infra.NewMock(t))

	cases := []struct {
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

	for _, tt := range cases {
		err := engine.ValidateConfs(tt.confs)
		assert.Equal(t, err, tt.err)
	}
}
