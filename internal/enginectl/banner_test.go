package enginectl

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func TestPrintBanner(t *testing.T) {
	cases := []struct {
		confs []model.Conf
		err error
		prepare func(*infra.Mock)
	}{
		{
			confs: []model.Conf{
				{ Host: "example.com", Origin: model.ConfOrigin{ Url: "localhost:3000" } },
				{ Host: "example2.com", Origin: model.ConfOrigin{ Url: "localhost:3001" } },
			},
			prepare: func(m *infra.Mock) {
				m.Log.EXPECT().Info("******************************")
				m.Log.EXPECT().Info("* Server started on localhost:3000")
				m.Log.EXPECT().Info("* ")
				m.Log.EXPECT().Info("* Sites:")
				m.Log.EXPECT().Info("* - %s (origin: %s)", "example.com", "localhost:3000")
				m.Log.EXPECT().Info("* - %s (origin: %s)", "example2.com", "localhost:3001")
				m.Log.EXPECT().Info("******************************")
			},
		},
	}

	for _, tt := range cases {
		engine := New()
		engine.Container = infra.NewMock(t, tt.prepare)

		engine.PrintBanner(tt.confs)
	}
}
