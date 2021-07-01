package main

import (
	"bytes"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

// curl -i --user "ming:mypwd" http://127.0.0.1:12345/static/

var user = []byte("ming")
var passwd = []byte("mypwd")

type ViewFunc func(http.ResponseWriter, *http.Request)

func Auth(f ViewFunc) ViewFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authPrefix := "Basic "

		// get request header
		auth := r.Header.Get("Authorization") // auth is []byte
		if strings.HasPrefix(auth, authPrefix) {
			// payload is []byte
			payload, err := base64.StdEncoding.DecodeString(
				auth[len(authPrefix):],
			)
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2) // pair is [][]byte
				if len(pair) == 2 && bytes.Equal(pair[0], user) &&
					bytes.Equal(pair[1], passwd) {
					f(w, r)
					return
				}
			}
		}
		// 401 Unauthorized
		// myrestricted can be other values (case sensitive)
		w.Header().Set("WWW-Authenticate", `Basic realm="myrestricted"`)
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func handleFileServer(dir, prefix string) ViewFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		realHandler(w, req)
	}
}

func main() {
	// file server, current directory as root
	http.HandleFunc("/static/", Auth(handleFileServer(".", "/static")))
	log.Fatal(http.ListenAndServe(":12345", nil))
}
