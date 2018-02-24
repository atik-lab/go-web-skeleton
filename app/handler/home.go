package controller

import(
	"github.com/atik-lab/go-web-skeleton/core"
	"net/http"
	"html/template"
	"github.com/atik-lab/go-web-skeleton/app"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.HiHandler(w, r)
}

func (h *Handler) More(w http.ResponseWriter, r *http.Request) {
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
		Version: core.GoWebIdentifier,
	}
	t, err := template.ParseFiles(h.config.Template + "/page.html")
	if err != nil { // if there is an error
		//log.Print("Template parsing error: ", err)
	}
	err = t.Execute(w, vars)
	if err != nil { // if there is an error
		//log.Print("Template executing error: ", err)
	}
}
