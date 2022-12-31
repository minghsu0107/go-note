package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// Fill it up
	ch <- 1

	select {
	case ch <- 2: // Put 2 in the channel unless it is full
	default:
		fmt.Println("Channel full. Discarding value")
	}
}
