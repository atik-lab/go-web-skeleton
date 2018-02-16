package main

import (
	"html/template"
	"net/http"
	"fmt"
	"strconv"
	"log"

	"github.com/gorilla/mux"
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

	// same in mux
	r := mux.NewRouter()

	// static files handler
	staticFiles := http.Dir(d.config.Static)
	staticFilesHandler := http.StripPrefix("/static/", http.FileServer(staticFiles))
	r.PathPrefix("/static/").Handler(staticFilesHandler).Methods("GET")

	// other
	r.HandleFunc("/{controller}/{action}/", d.preHandler(d.handler))

	var err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println(err)
	}
}

// Runs before handling
func (d *Daemon) preHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// Handler: Home
func (d *Daemon) handler(w http.ResponseWriter, r *http.Request) {
	/*
	varsm := mux.Vars(r)
	controller := varsm["controller"]
	action := varsm["action"]
	request := varsm["request"]
	*/

	// variables
	vars := struct {
		Version string
	}{
		Version: goWebIdentifier,
	}
	t, err := template.ParseFiles(d.config.Template + "/page.html")
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, vars)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err)
	}
}
