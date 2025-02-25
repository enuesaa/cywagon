package enginectl

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
)

func TestCalcMaxWaitForHealthy(t *testing.T) {
	cases := []struct {
		confs []model.Conf
		expect int
	}{
		{
			confs: []model.Conf{
				{
					Origin: model.ConfOrigin{
						WaitForHealthy: 50,
					},
				},
				{
					Origin: model.ConfOrigin{
						WaitForHealthy: 5,
					},
				},
				{
					Origin: model.ConfOrigin{
						WaitForHealthy: 1,
					},
				},
			},
			expect: 50,
		},
		{
			confs: []model.Conf{},
			expect: 1,
		},
	}

	for _, tt := range cases {
		engine := New()
		engine.Container = infra.NewMock(t)

		assert.Equal(t, engine.calcMaxWaitForHealthy(tt.confs), tt.expect)
	}
}
