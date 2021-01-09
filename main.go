package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/h00s-go/seismic-mqtt/config"
	"github.com/h00s-go/seismic-mqtt/seismic"
)

func main() {
	_, err := config.Load("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	s := seismic.New()
	defer s.Disconnect()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	events := make(chan seismic.Event)
	s.ReadMessages(events)

	for {
		select {
		case e := <-events:
			log.Println(e)
		case <-interrupt:
			return
		}
	}
}
