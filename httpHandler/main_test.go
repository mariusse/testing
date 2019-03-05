package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", index)
	return router
}

func TestIndex(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	router().ServeHTTP(response, request)
	if response.Body.String() != "Hello, World" {
		t.Errorf("\nexpected: Hello, World\ngot: %s", response.Body.String())
	}
}
