package main

import (
	"log"

	"github.com/BitMancers/Portigo/internal/engine"
	"github.com/BitMancers/Portigo/internal/resource"
)

func main() {
	conn, err := resource.InitLibvirt()

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	engine.Init(conn)
	engine.Start()

}
