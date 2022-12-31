package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	// handle only /hello, does not include /hello/
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", echoHandler)

	http.ListenAndServe(":8090", mux)
}
