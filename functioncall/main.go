package main

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

func main() {
	l := lua.NewState()
	error := lua.DoFile(l, "lib.lua")
	if error != nil {
		fmt.Printf("Error reading lib.lua: %s", error)
	}

	const n = 10
	l.Global("Fibonacci")
	l.PushInteger(n)
	l.Call(1, 1)
	result := [n]int{}
	for i := 0; i < n; i++ {
		l.RawGetInt(-1, i+1)
		result[i], _ = l.ToInteger(-1)
		l.Pop(1)
	}

	fmt.Printf("%v\n", result)
}
