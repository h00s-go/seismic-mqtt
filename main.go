package main

import (
	"log"

	"github.com/h00s-go/seismic-mqtt/config"
)

func main() {
	_, err := config.Load("config.toml")
	if err != nil {
		log.Fatal(err)
	}
}
