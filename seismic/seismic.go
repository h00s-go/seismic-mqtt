package seismic

// Seismic struct is object for communication with websocket
type Seismic struct {
}

// New creates new Seismic object
func New() (*Seismic, error) {
	return new(Seismic), nil
}
