package service

import (
	"fmt"
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
)

func TestConfServiceListConfPaths(t *testing.T) {
	cases := []struct {
		paths []string
		expected []string
		err error
		prepare func(*infra.Mock)
	}{
		{
			paths: []string{"aa.lua", "bb.lua"},
			expected: []string{"aa.lua", "bb.lua"},
			prepare: func(m *infra.Mock) {
				m.Fs.EXPECT().IsExist("aa.lua").Return(true)
				m.Fs.EXPECT().IsFile("aa.lua").Return(true)
				m.Fs.EXPECT().IsExist("bb.lua").Return(true)
				m.Fs.EXPECT().IsFile("bb.lua").Return(true)
			},
		},
		{
			paths: []string{"aa.lua", "sites"},
			expected: []string{"aa.lua", "sites/s-aa.lua", "sites/s-bb.lua"},
			prepare: func(m *infra.Mock) {
				m.Fs.EXPECT().IsExist("aa.lua").Return(true)
				m.Fs.EXPECT().IsFile("aa.lua").Return(true)
				m.Fs.EXPECT().IsExist("sites").Return(true)
				m.Fs.EXPECT().IsFile("sites").Return(false)
				m.Fs.EXPECT().ListFiles("sites").Return([]string{"sites/s-aa.lua", "sites/s-bb.lua", "sites/a.txt"}, nil)
			},
		},
		{
			paths: []string{"not-found"},
			err: fmt.Errorf("path not found: not-found"),
			prepare: func(m *infra.Mock) {
				m.Fs.EXPECT().IsExist("not-found").Return(false)
			},
		},
	}

	for _, tt := range cases {
		confsrv := ConfSrv{}
		confsrv.Container = infra.NewMock(t, tt.prepare)

		confpaths, err := confsrv.ListConfPaths(tt.paths)
		assert.Equal(t, err, tt.err)
		assert.Equal(t, confpaths, tt.expected)
	}
}

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
			},
		},
	}

	for _, tt := range cases {
		confsrv := ConfSrv{}
		confsrv.Container = infra.NewMock(t)

		err := confsrv.Validate(tt.conf)
		assert.Equal(t, err, tt.err)
	}
}
