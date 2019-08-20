// +build js,wasm

package main

import (
	"log"
	"syscall/js"
)

// TODO: log seems to cause an issue
func addition(this js.Value, args []js.Value) interface{} {
	log.Println("In WASM", args)
	a, b := args[0].Int(), args[1].Int()
	return a + b
}

func main() {
	ch := make(chan bool)

	// register functions
	fun := js.FuncOf(addition)
	js.Global().Set("addition", fun)

	res := js.Global().Get("proxy").Invoke(1, 2)
	log.Printf("1 + 2 = %d\n", res.Int())
	<-ch
}
