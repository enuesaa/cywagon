package enginectl

// import (
// 	"testing"

// 	"github.com/enuesaa/cywagon/internal/infra"
// 	"github.com/enuesaa/cywagon/internal/service/model"
// )

// func TestPrintBanner(t *testing.T) {
// 	cases := []struct {
// 		confs []model.Conf
// 		err error
// 		prepare func(*infra.Mock)
// 	}{
// 		{
// 			confs: []model.Conf{
// 				{ Host: "example.com" },
// 				{ Host: "example2.com" },
// 			},
// 			prepare: func(m *infra.Mock) {
// 				m.Log.EXPECT().Info("******************************")
// 				m.Log.EXPECT().Info("* Server started on localhost:3000")
// 				m.Log.EXPECT().Info("* ")
// 				m.Log.EXPECT().Info("* Sites:")
// 				m.Log.EXPECT().Info("* - %s", "example.com")
// 				m.Log.EXPECT().Info("* - %s", "example2.com")
// 				m.Log.EXPECT().Info("******************************")
// 			},
// 		},
// 	}

// 	for _, tt := range cases {
// 		engine := New()
// 		engine.Container = infra.NewMock(t, tt.prepare)

// 		engine.PrintBanner(tt.confs)
// 	}
// }
