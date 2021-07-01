package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Wait for them to finish (for a more robust approach, use a WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")
}

/*
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done
*/
/*
// return after get the first result and cancel remaining goroutines
// we can do this by closing a channel,
// because a receive operation on a closed channel can always proceed immediately,
// yielding the element type's zero value
// if we don't do this, the remaining goroutines will block forever
// because the reciever of c already exits
func First(query string, replicas ...Search) Result {
    c := make(chan Result)
    done := make(chan struct{})
    defer close(done) // unblock all the senders
    searchReplica := func(i int) {
        select {
        case c <- replicas[i](query):
        case <-done:
            return
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }

    return <-c
}
*/
