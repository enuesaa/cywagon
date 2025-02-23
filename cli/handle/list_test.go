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
		prepare func(*infra.MockFsInterface)
	}{
		{
			paths: []string{"aa.lua", "bb.lua"},
			expected: []string{"aa.lua", "bb.lua"},
			prepare: func(fs *infra.MockFsInterface) {
				fs.EXPECT().IsDir("aa.lua").Return(false, nil)
				fs.EXPECT().IsDir("bb.lua").Return(false, nil)
			},
		},
	}

	for _, tt := range cases {
		handler := New()
		handler.Container = infra.NewMock(t)
	
		tt.prepare(handler.Container.Fs.(*infra.MockFsInterface))
	
		confpaths, err := handler.listConfPaths(tt.paths)
		require.Nil(t, err)
		assert.Equal(t, confpaths, tt.expected)
	}
}
