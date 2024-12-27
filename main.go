package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yuin/gopher-lua"
)

func main() {
	luaScript := `
		if string.find(event.path, "7") then
			print("bbb")
		end
		event.path = event.path .. "aaa"
	`
	L := lua.NewState()
	defer L.Close()

	start := time.Now()

	for range 10000 {
		event := L.NewTable()
		L.SetField(event, "path", lua.LString(fmt.Sprintf("example%d/", rand.Intn(1000))))
		L.SetGlobal("event", event)
		if err := L.DoString(luaScript); err != nil {
			panic(err)
		}
		updatedObj := L.GetGlobal("event").(*lua.LTable)
		path := L.GetField(updatedObj, "path").(lua.LString)
		fmt.Printf("path: %s\n", path)
	}

	elapsed := time.Since(start)
	fmt.Printf("time: %v\n", elapsed)
}

