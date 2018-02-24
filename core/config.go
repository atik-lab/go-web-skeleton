package core

type Config struct {
	Debug			bool
	Port			uint
	Verbose			bool
	Template		string
	Static			string
}

// Get config with defaults
func NewConfig() *Config {
	return &Config{
		Debug: false,
		Port: 8080,
		Verbose: true,
		Template: "app/template/",
		Static: "static/",
	}
}
