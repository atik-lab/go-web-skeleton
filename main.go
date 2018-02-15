package main

import (
	"flag"
)

var (
	config		*Config
	daemon		*Daemon
)

// Main, get flags and start daemon
func main() {
	config = NewConfig()

	// Parse flags and process config
	flag.UintVar(&config.Port, "p", config.Port, "Port to listen")
	flag.BoolVar(&config.Verbose, "v", config.Verbose, "Verbose, show output")
	flag.BoolVar(&config.Debug, "d", config.Debug, "Debug")
	flag.Parse()

	// Start daemon
	daemon = NewDaemon(config)
	daemon.Start()
}
