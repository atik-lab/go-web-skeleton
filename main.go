package main

import (
	"flag"
	"github.com/atik-lab/go-web-skeleton/core"
)

var (
	config		*core.Config
	daemon		*core.Daemon
)

// Main, get flags and start daemon
func main() {
	config = core.NewConfig()

	// Parse flags and process config
	flag.UintVar(&config.Port, "p", config.Port, "Port to listen")
	flag.BoolVar(&config.Verbose, "v", config.Verbose, "Verbose, show output")
	flag.BoolVar(&config.Debug, "d", config.Debug, "Debug")
	flag.Parse()

	// Start daemon
	daemon = core.NewDaemon(config)
	daemon.Start()
}
