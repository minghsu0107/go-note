package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "result 1"
	}()

	// Since select proceeds with the first receive that’s ready,
	// we’ll take the timeout case if the operation takes more than the allowed 3s
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 2")
	}
}

/*
timeout 1
result 2
*/
