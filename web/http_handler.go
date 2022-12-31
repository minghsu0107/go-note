package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/smth", smthHandler)
	http.Handle("/oldsmth", http.RedirectHandler("/smth", http.StatusMovedPermanently))

	http.ListenAndServe(":8090", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// only handle "/"
	// ex: localhost:8090/a will be rounted to this handler and causing error
	if r.URL.Path != "/" {
		// http.NotFound(w, req)
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome home")
}

func smthHandler(w http.ResponseWriter, r *http.Request) {
	// only handle "/smth"
	// ex: localhost:8090/smth/a will be rounted to this handler and causing error
	if r.URL.Path != "/smth" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome smth")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
