package handle

import (
	"fmt"
	"testing"

	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCheck(t *testing.T) {
	cases := []struct {
		paths []string
		err error
		prepareEngine func(*enginectl.MockEngineInterface)
		prepareConfSrv func(*service.MockConfSrvInterface)
	}{
		{
			paths: []string{},
			prepareEngine: func(e *enginectl.MockEngineInterface) {
				confs := []model.Conf{
					{ Host: "example.com" },
				}
				e.EXPECT().ValidateConfs(confs).Return(nil)
			},
			prepareConfSrv: func(s *service.MockConfSrvInterface) {
				confs := []model.Conf{
					{ Host: "example.com" },
				}
				s.EXPECT().List(gomock.Any()).Return(confs, nil)
			},
		},
		{
			paths: []string{},
			prepareEngine: func(e *enginectl.MockEngineInterface) {
				confs := []model.Conf{
					{ Host: "example.com" },
				}
				err := fmt.Errorf("conf err")
				e.EXPECT().ValidateConfs(confs).Return(err)
			},
			prepareConfSrv: func(s *service.MockConfSrvInterface) {
				confs := []model.Conf{
					{ Host: "example.com" },
				}
				s.EXPECT().List(gomock.Any()).Return(confs, nil)
			},
			err: fmt.Errorf("conf err"),
		},
	}

	for _, tt := range cases {
		handler := New()
		handler.Engine = enginectl.NewMock(t, tt.prepareEngine)
		handler.ConfSrv = service.NewConfSrvMock(t, tt.prepareConfSrv)

		err := handler.Check(tt.paths)
		assert.Equal(t, err, tt.err)
	}
}
