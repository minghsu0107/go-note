package main

import (
	"fmt"
	"net/http"
	"sync"
)

var table sync.Map

type Getter func() *http.Request

func GetRequest(url string) *http.Request {
	getter := getReqFromMap(url)
	return getter()
}

func getReqFromMap(url string) Getter {
	if f, ok := table.Load(url); ok {
		return f.(Getter)
	}

	var req *http.Request
	var once sync.Once
	wrapGetter := Getter(func() *http.Request {
		once.Do(func() {
			req, _ = http.NewRequest(http.MethodGet, url, nil)
			fmt.Printf("Done: %s\n", url)
		})

		return req
	})

	f, loaded := table.LoadOrStore(url, wrapGetter)
	// already stored by another concurrent goroutine
	if loaded {
		return f.(Getter)
	}

	return wrapGetter
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if i%2 == 0 {
				GetRequest("http://example.com/user1/")
			} else {
				GetRequest("http://example.com/user2/")
			}
		}(i)
	}
	wg.Wait()
}
