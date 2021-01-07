package seismic

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

const urlHost = "www.seismicportal.eu"
const urlPath = "/standing_order/websocket"
const pongWait = 60 * time.Second

// Seismic struct is object for communication with websocket
type Seismic struct {
	conn *websocket.Conn
}

// New creates new Seismic object
func New() *Seismic {
	return new(Seismic)
}

// Connect connects to Seismic portal websocket
func (s *Seismic) Connect() error {
	u := url.URL{Scheme: "wss", Host: urlHost, Path: urlPath}

	var err error
	s.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	s.conn.SetPongHandler(func(string) error { s.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	return nil
}

// Disconnect disconnects from Seismic portal websocket
func (s *Seismic) Disconnect() error {
	return s.conn.Close()
}

// ReadMessages reads new events (json) from seismic portal, parse it and sends to channel
func (s *Seismic) ReadMessages(e chan Event) {
	go func() {
		for {
			_, message, err := s.conn.ReadMessage()
			if err == nil {
				if event, err := ParseEvent(message); err == nil {
					e <- event
				}
			} else {
				log.Println(err)
				s.Disconnect()
				s.Connect()
			}
		}
	}()
}
