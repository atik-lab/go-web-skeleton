package main

import (
	"html/template"
	"net/http"
	"fmt"
	"strconv"
	"log"
	"path/filepath"
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
	// variables
	var port string = ":" + strconv.Itoa(int(d.config.Port))
	// static files
	filePrefix, _ := filepath.Abs("static")
	fs := http.FileServer(http.Dir(filePrefix))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// handlers
	http.HandleFunc("/", d.handlerHome)
	// start server
	var err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// Handler: Home
func (d *Daemon) handlerHome(w http.ResponseWriter, r *http.Request) {
	// variables
	vars := struct {
		Version string
	}{
		Version: goWebIdentifier,
	}
	filePrefix, _ := filepath.Abs(d.config.Design)
	t, err := template.ParseFiles(filePrefix + "/page.html")
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, vars)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err)
	}
}
