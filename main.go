package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/h00s-go/seismic-mqtt/config"
	"github.com/h00s/goseismic"
)

func main() {
	_, err := config.Load("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	s := goseismic.NewSeismic()
	s.Connect()
	defer s.Disconnect()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case e := <-s.Events:
			log.Println(e)
		case <-interrupt:
			return
		}
	}
}
