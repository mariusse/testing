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
		t.Errorf("\nexpected %s, got %s", "Hello, World", response.Body.String())
	}
}
