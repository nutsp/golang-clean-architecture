package main

import "github.com/nutsp/golang-clean-architecture/internal/container"

func main() {
	cn := container.NewContainer()
	if err := cn.Run(); err != nil {
		panic(err)
	}
}
