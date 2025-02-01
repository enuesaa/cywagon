package liblua

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	lua "github.com/yuin/gopher-lua"
)

func TestUnmarshal(t *testing.T) {
	code := `
	entry = {}
	entry.a = "aaa"
	entry.b = 1
	`
	type Entry struct {
		A string `lua:"a"`
		B int `lua:"b"`
	}
	var entry Entry

	state := lua.NewState()
	err := state.DoString(code)
	require.Nil(t, err)

	table := state.GetGlobal("entry").(*lua.LTable)
	err = Unmarshal(table, &entry)
	require.Nil(t, err)

	assert.Equal(t, "aaa", entry.A)
	assert.Equal(t, 1, entry.B)
}
