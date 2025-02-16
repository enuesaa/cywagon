package service

import (
	"testing"

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

	for _, tt := range table {
		confsrv := NewConfService()
		err := confsrv.Validate(tt.conf)
		assert.Equal(t, err, tt.err)
	}
}
