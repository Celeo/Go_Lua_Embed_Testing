package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

// output prints "engine messages" to the console
func output(s string, a ...interface{}) {
	fmt.Printf(">>> "+s+"\n", a...)
}

// Add adds two numbers together and returns the new value.
func Add(L *lua.LState) int {
	n1 := L.ToNumber(1)
	n2 := L.ToNumber(2)
	added := doAdd(n1, n2)
	output("Read '%d' and '%d' from Lua calling Add() -> result is '%d'", n1, n2, added)
	L.Push(added)
	return 1
}

// doAdd adds two Lua Numbers together
func doAdd(n1 lua.LNumber, n2 lua.LNumber) lua.LNumber {
	return n1 + n2
}

// Cap returns the first number if it's under the second, otherwise the second.
// The return value is the "cap" of the 2 arguments.
func Cap(L *lua.LState) int {
	n1 := L.ToNumber(1)
	n2 := L.ToNumber(2)
	if n1 < n2 {
		L.Push(lua.LNumber(n1))
	} else {
		L.Push(lua.LNumber(n2))
	}
	return 1
}
