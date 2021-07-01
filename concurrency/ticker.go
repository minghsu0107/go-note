package main

import (
	"fmt"
	"time"
)

// tickers are for when you want to do something repeatedly at regular intervals.

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

/*
Tick at 2020-05-04 17:48:12.499161 +0800 CST m=+0.501623312
Tick at 2020-05-04 17:48:13.002595 +0800 CST m=+1.005054108
Tick at 2020-05-04 17:48:13.501283 +0800 CST m=+1.503737673
Ticker stopped
*/
