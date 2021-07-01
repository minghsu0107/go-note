package main

import (
	"fmt"
	"time"
)

func main() {
	data := []string{"one", "two", "three"}

	// the closure function's v is always the same one
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second) // three three three

	for _, v := range data {
		vCopy := v
		go func() {
			fmt.Println(vCopy)
		}()
	}
	time.Sleep(3 * time.Second) // one two three

	for _, v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v)
	}
	time.Sleep(3 * time.Second) // one two three
}
