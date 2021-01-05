package main

import (
	"log"

	"github.com/h00s-go/seismic-mqtt/config"
	"github.com/h00s-go/seismic-mqtt/seismic"
)

func main() {
	_, err := config.Load("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	_, err = seismic.New()
	if err != nil {
		log.Fatal(err)
	}
}
