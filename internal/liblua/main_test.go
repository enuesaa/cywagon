package liblua

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	lua "github.com/yuin/gopher-lua"
)

func NewForTesting(t *testing.T) Runner {
	return Runner{
		Container: infra.NewMock(t),
		state: lua.NewState(),
	}
}
