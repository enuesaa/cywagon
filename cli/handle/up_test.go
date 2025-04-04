package handle

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUp(t *testing.T) {
	cases := []struct {
		paths []string
		err error
		prepareContainer func(*infra.Mock)
		prepareEngine func(*enginectl.MockEngineInterface)
		prepareConfSrv func(*service.MockConfSrvInterface)
	}{
		{
			paths: []string{},
			prepareContainer: func(m *infra.Mock) {
				m.Log.EXPECT().Info("Start up sites..")
				m.Log.EXPECT().Info("Start health check..")
			},
			prepareEngine: func(e *enginectl.MockEngineInterface) {
				e.EXPECT().PrintBanner(gomock.Any()).Return()
				e.EXPECT().Serve(gomock.Any()).Return(nil)
			},
			prepareConfSrv: func(s *service.MockConfSrvInterface) {
				confs := []model.Conf{
					{ Host: "example.com" },
				}
				s.EXPECT().List(gomock.Any()).Return(confs, nil)
			},
		},
	}

	for _, tt := range cases {
		handler := New()
		handler.Container = infra.NewMock(t, tt.prepareContainer)
		handler.Engine = enginectl.NewMock(t, tt.prepareEngine)
		handler.ConfSrv = service.NewConfSrvMock(t, tt.prepareConfSrv)

		err := handler.Up(tt.paths)
		assert.Equal(t, err, tt.err)
	}
}
