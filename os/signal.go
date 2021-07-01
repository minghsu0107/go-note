package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	ticker := time.NewTicker(500 * time.Millisecond)

	// registers the given channel to receive notifications of the specified signals
	signal.Notify(sigs, []os.Signal{syscall.SIGINT, syscall.SIGTERM}...)

	run := func(id int) {
		for {
			select {
			case sig := <-sigs:
				fmt.Println()
				fmt.Println(sig)
				done <- true
				return
			case <-ticker.C:
				fmt.Println("working ", id)
			}
		}
	}

	go run(1)
	go run(2)

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
