package config

import "github.com/BurntSushi/toml"

// Configuration struct have all fields from configuration TOML file
type Configuration struct {
	Log Log
}

// Log defines logging configuration (log filename)
type Log struct {
	Filename string
}

// Load loads configuration from path
func Load(path string) (*Configuration, error) {
	c := new(Configuration)

	if _, err := toml.DecodeFile(path, c); err != nil {
		return c, err
	}

	return c, nil
}
