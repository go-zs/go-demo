package main

import (
	"github.com/go-zs/go-demo/di/digdemo/di"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(server *di.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}

}
