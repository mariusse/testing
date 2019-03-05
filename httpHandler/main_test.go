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

// inspirert fra nett. Er den riktig? Må/burde vel kalle response.Result()?
// noe voldsomt med mux og kanskje?
func TestIndex(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)
	if w.Body.String() != "Hello, World" {
		t.Errorf("\nexpected: Hello, World\ngot: %s", w.Body.String())
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

// funker like herlig uten w.Result() forsåvidt...
func TestIndex3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	index(w, req)
	// resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)
	if string(w.Body.String()) != "Hello, World" {
		t.Errorf("\nexpected: Hello, World\ngot: %s", w.Body.String())
	}
}