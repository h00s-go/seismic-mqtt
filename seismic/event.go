package seismic

import (
	"encoding/json"
	"log"
	"time"
)

// Event represent one information obtained from seismic portal
type Event struct {
	Data struct {
		Properties struct {
			Region    string    `json:"flynn_region"`
			Magnitude float64   `json:"mag"`
			Depth     float64   `json:"depth"`
			Latitude  float64   `json:"lat"`
			Longitude float64   `json:"lon"`
			UpdatedAt time.Time `json:"lastupdate"`
		} `json:"properties" binding:"required"`
	} `json:"data" binding:"required"`
}

// ParseEvent converts bytes to Event struct
func ParseEvent(data []byte) (Event, error) {
	var event Event

	if err := json.Unmarshal(data, &event); err != nil {
		log.Println("Error parsing JSON")
		return event, err
	}

	return event, nil
}
