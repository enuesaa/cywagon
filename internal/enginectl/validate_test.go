package enginectl

import (
	"fmt"
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	engine := Engine{
		Container: infra.NewMock(t),
	}

	table := []struct{
		confs []model.Conf
		err   error
	}{
		{
			confs: []model.Conf{
				{
					Host: "",
				},
			},
			err: fmt.Errorf("host is required"),
		},
		{
			confs: []model.Conf{
				{
					Host: "example.com",
				},
			},
			err: nil,
		},
	}

	for _, tt := range table {
		err := engine.Validate(tt.confs)
		assert.Equal(t, err, tt.err)
	}
}
