package controller

import (
	"net/http"
	"encoding/json"
	"github.com/atik-lab/go-web-skeleton/core"
)

type Handler struct {
	config 	*core.Config
}

// Create
func NewHandler(config *core.Config) *Handler {
	return &Handler {
		config,
	}
}

// Handler: Hi
func (h *Handler) HiHandler(w http.ResponseWriter, r *http.Request) {
	h.WriteJSON(w, struct {
		string
	} {
		"Hi!",
	})
}

// Ping Handler
func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	h.WriteSuccess(w)
}

// Write Json, used in handlers to send messages to the connected peer
func (h *Handler) WriteJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(obj)
	if _, isJsonErr := err.(*json.SyntaxError); isJsonErr {
		// FAIL
	}
}

// Write success, used in handlers to send success response to the connected peer
func (h *Handler) WriteSuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
