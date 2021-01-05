package seismic

import "time"

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
