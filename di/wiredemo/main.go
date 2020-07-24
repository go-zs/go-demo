package main

import "github.com/go-zs/go-demo/di/wiredemo/wire"

// 使用wire前
/*
func main() {
	message := constructor.NewMessage("hello world")
	greeter := constructor.NewGreeter(message)
	event := constructor.NewEvent(greeter)

	event.Start()
}
*/

// 使用wire后
func main() {
	event := wire.InitializeEvent("hello_world")

	event.Start()
}
