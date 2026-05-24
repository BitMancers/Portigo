package main

import (
	"portigo/internal/engine"
	"portigo/internal/resource"
)

func main() {
	resource.InitLibvirt()
	engine.Init()
	engine.Start()
}
