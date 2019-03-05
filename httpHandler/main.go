package main

import "net/http"

func main() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello, World"))
}
