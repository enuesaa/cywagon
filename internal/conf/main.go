package confctl

import (
	"github.com/yuin/gopher-lua"
)

var ExampleScript = `
	if string.find(event.path, "7") then
		print("bbb")
	end
	event.path = event.path .. "aaa"
`

func Run(script string) string {
	L := lua.NewState()
	defer L.Close()

	event := L.NewTable()
	L.SetField(event, "path", lua.LString("example/"))
	L.SetGlobal("event", event)

	if err := L.DoString(script); err != nil {
		panic(err)
	}
	updatedObj := L.GetGlobal("event").(*lua.LTable)
	path := L.GetField(updatedObj, "path").(lua.LString)

	return string(path)
}
