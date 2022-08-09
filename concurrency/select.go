package main

import (
	"fmt"
	"time"
)

// Go’s select lets you wait on multiple channel operations

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		close(c0)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	select {
	case <-c0:
		fmt.Println("read on closed channel")
	}

	// will not enter this loop
	for i := range c0 {
		fmt.Println("will not enter this loop")
		fmt.Println(i)
	}

	// use select to await both of these values simultaneously,
	// printing each one as it arrives
	// if both cases happen at the same time, go will choose one randomly!
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	// Note that the total execution time is only ~2 seconds
	// since both the 1 and 2 second Sleeps execute concurrently
}
