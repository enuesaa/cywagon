package service

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
)

func TestConfServiceValidate(t *testing.T) {
	table := []struct{
		conf model.Conf
		err  error
	}{
		{
			conf: model.Conf{
				Host: "",
			},
			err: ErrConfHostRequired,
		},
		{
			conf: model.Conf{
				Host: "example.com",
				Entry: model.ConfEntry{
					Cmd: "go run .",
				},
			},
		},
		{
			conf: model.Conf{
				Host: "example.com",
				Entry: model.ConfEntry{
					Cmd: "",
				},
			},
			err: ErrConfEntryCmdRequired,
		},
	}

	confsrv := ConfService{
		Container: infra.NewMock(t),
	}
	for _, tt := range table {
		err := confsrv.Validate(tt.conf)
		assert.Equal(t, err, tt.err)
	}
}
