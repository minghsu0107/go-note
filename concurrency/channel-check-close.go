package main

import "fmt"

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		// read on close channel
		return true
	default:
	}

	return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}
