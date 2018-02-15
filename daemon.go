package main

import (
	"net/http"
	"fmt"
	"strconv"
)

const (
	goWebIdentifier		= "go-web"	// node identifier to show
	goWebVersion		= "0.0.1"	// node version to show
)

type Daemon struct {
	config		*Config
}

// How to create a new Daemon
func NewDaemon(config *Config) *Daemon {
	return &Daemon{
		config,
	}
}

// Start daemon
func (d *Daemon) Start() {
	var port string = ":" + strconv.Itoa(int(d.config.Port))
	http.HandleFunc("/", helloWorld)
	var err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// hello world
func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, " Hello World")
}
