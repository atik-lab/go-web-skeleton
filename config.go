package main

type Config struct {
	Debug			bool
	Port			uint
	Verbose			bool
	Design			string
}

// Get config with defaults
func NewConfig() *Config {
	return &Config{
		Debug: false,
		Port: 8080,
		Verbose: true,
		Design: "design",
	}
}
