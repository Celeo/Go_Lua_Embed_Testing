package main

import (
	"github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("engine_add", L.NewFunction(Add))

	if err := L.DoFile("script.lua"); err != nil {
		panic(err)
	}
}
