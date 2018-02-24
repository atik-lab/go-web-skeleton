package core

import (
	"net/http"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/atik-lab/go-web-skeleton/app"
)

const (
	GoWebIdentifier		= "go-web"	// node identifier to show
	GoWebVersion		= "0.0.1"	// node version to show
)

type Daemon struct {
	config		*Config
	logger		*Logger
}

// How to create a new Daemon
func NewDaemon(config *Config) *Daemon {
	return &Daemon{
		config,
		NewLogger("goweb.log", true, LevelInfo, "daemon"),
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
	var router = app.NewRouter()
	router.Route()

	var err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println(err)
	}
}

// Runs before handling
func (d *Daemon) preHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d.logger.Log(r.URL.Path, LevelInfo)
		f(w, r)
	}
}

