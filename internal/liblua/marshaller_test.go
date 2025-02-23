package liblua

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	lua "github.com/yuin/gopher-lua"
)

func TestMarshal(t *testing.T) {
	runner := New()
	runner.Container = infra.NewMock(t).Container()

	type Entry struct {
		A string `lua:"a"`
		B int    `lua:"b"`
	}
	entry := Entry{
		A: "aaa",
		B: 1,
	}
	table, err := runner.Marshal(entry)
	require.Nil(t, err)

	state := lua.NewState()
	assert.Equal(t, "aaa", string(state.GetField(table, "a").(lua.LString)))
	assert.Equal(t, 1, int(state.GetField(table, "b").(lua.LNumber)))
}

func TestUnmarshal(t *testing.T) {
	runner := New()
	runner.Container = infra.NewMock(t).Container()

	code := `
	entry = {}
	entry.a = "aaa"
	entry.b = 1
	`
	type Entry struct {
		A string `lua:"a"`
		B int    `lua:"b"`
	}
	var entry Entry

	state := lua.NewState()
	err := state.DoString(code)
	require.Nil(t, err)

	table := state.GetGlobal("entry").(*lua.LTable)
	err = runner.Unmarshal(table, &entry)
	require.Nil(t, err)

	assert.Equal(t, "aaa", entry.A)
	assert.Equal(t, 1, entry.B)
}
