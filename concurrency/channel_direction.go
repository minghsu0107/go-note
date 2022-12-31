package main

import "fmt"

// you can specify if a channel is meant to only send or receive values
// chan<- type: the channel receives msg only
// <-chan type: the channel sends msg only
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs) // passed message
}
