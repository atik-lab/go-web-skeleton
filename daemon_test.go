package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestHandler(t *testing.T) {
	// make a new request
	req, err := http.NewRequest("GET", "", nil)

	// if error
	if err != nil {
		t.Fatal(err)
	}

	// this recorder will act as the target of our http request
	recorder := httptest.NewRecorder()

	// Create an HTTP handler
	d := NewDaemon(NewConfig())
	hf := http.HandlerFunc(d.handler)
	// Serve the HTTP request to our recorder
	hf.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `book`
	actual := recorder.Body.String()
	if strings.Index(actual, expected) == -1 {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}