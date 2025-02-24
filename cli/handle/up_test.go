package handle

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUp(t *testing.T) {

	cases := []struct {
		paths []string
		err error
		prepareContainer func(*infra.Mock)
		prepareEngine func(*enginectl.MockEngineCtl)
	}{
		{
			paths: []string{"aa.lua", "bb.lua"},
			prepareContainer: func(m *infra.Mock) {
				m.Fs.EXPECT().IsExist("aa.lua").Return(true)
				m.Fs.EXPECT().IsFile("aa.lua").Return(true)
				m.Fs.EXPECT().IsExist("bb.lua").Return(true)
				m.Fs.EXPECT().IsFile("bb.lua").Return(true)
			},
			prepareEngine: func(e *enginectl.MockEngineCtl) {
				e.EXPECT().StartUp(gomock.Any()).Return(nil)
				e.EXPECT().StartHealthCheck(gomock.Any()).Return(nil)
				e.EXPECT().PrintBanner(gomock.Any()).Return()
				e.EXPECT().Serve(gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range cases {
		container := infra.NewMock(t, tt.prepareContainer)

		handler := New(container)
		handler.Engine = enginectl.NewMock(t, tt.prepareEngine)

		err := handler.Up(tt.paths)
		assert.Equal(t, err, tt.err)
	}
}
