package handle

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListConfPaths(t *testing.T) {
	cases := []struct {
		paths []string
		expected []string
		prepare func(*infra.Mock)
	}{
		{
			paths: []string{"aa.lua", "bb.lua"},
			expected: []string{"aa.lua", "bb.lua"},
			prepare: func(m *infra.Mock) {
				m.Fs.EXPECT().IsDir("aa.lua").Return(false, nil)
				m.Fs.EXPECT().IsDir("bb.lua").Return(false, nil)
			},
		},
		{
			paths: []string{"aa.lua", "sites"},
			expected: []string{"aa.lua", "sites/s-aa.lua", "sites/s-bb.lua"},
			prepare: func(m *infra.Mock) {
				m.Fs.EXPECT().IsDir("aa.lua").Return(false, nil)
				m.Fs.EXPECT().IsDir("sites").Return(true, nil)
				m.Fs.EXPECT().ListFiles("sites").Return([]string{"sites/s-aa.lua", "sites/s-bb.lua", "sites/a.txt"}, nil)
			},
		},
	}

	for _, tt := range cases {
		handler := New()
		handler.Container = infra.NewMock(t, tt.prepare)

		confpaths, err := handler.listConfPaths(tt.paths)
		require.Nil(t, err)
		assert.Equal(t, confpaths, tt.expected)
	}
}
