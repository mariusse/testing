package main

import (
	"io/ioutil"
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

// er denne riktigere?
func TestIndex2(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	index(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Hello, World" {
		t.Errorf("\nexpected: Hello, World\ngot: %s", string(body))
	}
}
