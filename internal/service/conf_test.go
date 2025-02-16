package service

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
)

func TestConfServiceValidate(t *testing.T) {
	cases := []struct {
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

	confsrv := NewConfService()
	confsrv.Container = infra.NewMock(t)

	for _, tt := range cases {
		err := confsrv.Validate(tt.conf)
		assert.Equal(t, err, tt.err)
	}
}
