package main

import lua "github.com/Shopify/go-lua"

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	l.Global("print")
	l.PushString("Hello world!")
	l.Call(1, 0)
}
