package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// first way
type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

// second way
func helloHandlerFunc(w http.ResponseWriter, req *http.Request) {
	var ip string = req.RemoteAddr
	log.Println(ip)
	// parse query string (get) or form value (post)
	if err := req.ParseForm(); err != nil {
		log.Printf("ParseForm() err: %v\n", err)
		return
	}
	user := req.FormValue("user")
	log.Println(user)
	io.WriteString(w, "hello, world\n")
}

// third way
func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	// h1 := req.Header.Get("Header1") // return the "first" value of header1

	// this should precede WriteHeader
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Hello-World", "*")
	w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	w.Header().Add("Trailer", "AtEnd3")
	// this should precede Write to w
	w.WriteHeader(http.StatusOK)

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	// any header will be forced to be canonical key in go
	// for example, curl localhost:8090/headers -H 'foo: 123'
	// it will become 'Foo: 123'
	// h2 := req.Header["foo"][0] // invalid
	// h2 := req.Header.get("foo") // valid
	// h2 := req.Header[textproto.CanonicalMIMEHeaderKey("foo")][0] // valid
	// note that we should import "net/textproto"
}

func main() {
	// background: go run http-server.go &
	hh := http.HandlerFunc(helloHandlerFunc)
	http.Handle("/hello", hh)

	http.HandleFunc("/hello2", hello)
	http.HandleFunc("/headers", headers)

	http.Handle("/", &helloHandler{}) // this matches all other url patterns
	// nil tells it to use the default router we’ve just set up
	log.Fatal(http.ListenAndServe(":8090", nil))
}
