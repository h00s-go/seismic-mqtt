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
	conn      *websocket.Conn
	connected bool
	Events    chan Event
}

// New creates new Seismic object
func New() *Seismic {
	return &Seismic{
		conn:      nil,
		connected: false,
		Events:    make(chan Event),
	}
}

// Connect connects to Seismic portal websocket
func (s *Seismic) Connect() {
	u := url.URL{Scheme: "wss", Host: urlHost, Path: urlPath}

	var err error
	for {
		log.Println("Trying to connect to websocket...")
		s.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
		if err == nil {
			s.connected = true
			s.conn.SetPongHandler(func(string) error { s.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
			s.conn.SetCloseHandler(func(int, string) error {
				log.Println("Closed connection to websocket.")
				s.connected = false
				return nil
			})
			log.Println("Connected to websocket!")
			break
		}
		time.Sleep(10 * time.Second)
	}
}

// Disconnect disconnects from Seismic portal websocket
func (s *Seismic) Disconnect() error {
	s.connected = false
	return s.conn.Close()
}

// ReadMessages reads new events (json) from seismic portal, parse it and sends to channel
func (s *Seismic) ReadMessages() {
	go func() {
		for {
			if !s.connected {
				s.Connect()
			}
			_, message, err := s.conn.ReadMessage()
			if err == nil {
				if event, err := ParseEvent(message); err == nil {
					s.Events <- event
				}
			} else {
				log.Println(err)
				s.Disconnect()
			}
		}
	}()
}
