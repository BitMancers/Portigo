package main

import (
	"log"

	"github.com/BitMancers/Portigo/internal/engine"
	"github.com/BitMancers/Portigo/internal/resource"
)

func main() {
	_, err := resource.InitLibvirt()
	if err != nil {
		log.Fatal(err)
	}
	engine.Init()
	engine.Start()
}
