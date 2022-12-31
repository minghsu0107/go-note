package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine
// and receive those values into another goroutine
// Send a value into a channel using the channel <- syntax
// The <-channel syntax receives a value from the channel

// example of using a blocking receive to wait for a goroutine to finish
func worker(done *chan bool) {
	fmt.Println("working...")
	time.Sleep(1 * time.Second)
	fmt.Println("done")

	*done <- true
}

func main() {

	messages := make(chan string)
	go func() { messages <- "ping" }()
	// By default sends and receives block until both the sender and receiver are ready
	msg := <-messages
	fmt.Println(msg) // ping

	// By default channels are unbuffered,
	// meaning that they will only accept sends (chan <-) i
	// if there is a corresponding receive (<- chan) ready to receive the sent value
	// Buffered channels accept a limited number of values
	// without a corresponding receiver for those values.
	// Note that the channel is buffered, so the send in the goroutine is nonblocking.
	messages2 := make(chan string, 2) // 2
	fmt.Println(cap(messages2))
	messages2 <- "buffered"
	messages2 <- "channel"
	fmt.Println(<-messages2) // buffered
	fmt.Println(<-messages2) // channel

	done := make(chan bool, 2)
	// pass by pointer
	go worker(&done)
	go worker(&done)
	// Block until we receive all notifications from the worker on the channel
	<-done
	<-done
}
