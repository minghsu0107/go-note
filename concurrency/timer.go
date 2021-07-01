package main

import (
	"fmt"
	"time"
)

// f you just wanted to wait, you could have used time.Sleep.
// One reason a timer may be useful is that you can cancel the timer before it fires

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

/*
Timer 1 fired
Timer 2 stopped
*/
