package main

import (
	"fmt"
	"time"
)

// Go’s select lets you wait on multiple channel operations

func main() {
	c0 := make(chan string)
	c1 := make(chan string, 1)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		close(c0)

		c1 <- "hello"
		close(c1)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "two"
	}()

	select {
	case <-c0:
		fmt.Println("read on closed buffered channel")
	}

	select {
	case v := <-c1:
		fmt.Printf("read value from closed buffered channel: %v\n", v)
	}

	// will not enter this loop
	for i := range c0 {
		fmt.Println("will not enter this loop")
		fmt.Println(i)
	}

	// will not enter this loop
	for i := range c1 {
		fmt.Println("will not enter this loop")
		fmt.Println(i)
	}

	// use select to await both of these values simultaneously,
	// printing each one as it arrives
	// if both cases happen at the same time, go will choose one randomly!
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c2:
			fmt.Println("received", msg1)
		case msg2 := <-c3:
			fmt.Println("received", msg2)
		}
	}
	// Note that the total execution time is only ~2 seconds
	// since both the 1 and 2 second Sleeps execute concurrently
}
