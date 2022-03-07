package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)
// errgroup returns the first error from a group of goroutines
func main() {
	eg := errgroup.Group{}
	eg.Go(func() error {
		return getPage("https://dne.minghsu.io")
	})
	eg.Go(func() error {
		return getPage("https://google.com")
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("get error: %v", err)
	}
}

func getPage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fail to get page: %s, wrong statusCode: %d", url, resp.StatusCode)
	}
	log.Printf("success get page %s", url)
	return nil
}
// 2022/03/07 22:56:14 success get page https://google.com
// 2022/03/07 22:56:14 get error: Get "https://dne.minghsu.io": dial tcp: lookup dne.minghsu.io: no such host
// exit status 1
