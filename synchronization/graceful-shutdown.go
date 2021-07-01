package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type atomicBool int32

func (b *atomicBool) isTrue() bool { return atomic.LoadInt32((*int32)(b)) != 0 }
func (b *atomicBool) setTrue()     { atomic.StoreInt32((*int32)(b), 1) }
func (b *atomicBool) setFalse()    { atomic.StoreInt32((*int32)(b), 0) }

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// capture the Ctrl-C os signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	var isEnd atomicBool
	isEnd.setFalse()

	// only one thread consumes the signal channel
	go func() {
		// we wait for the Ctrl-C signal here
		<-signalChan
		isEnd.setTrue()
		fmt.Println("signal reciever thread exited")
	}()

	for {
		if isEnd.isTrue() {
			fmt.Println("stopped")
			break
		}
		t := <-ticker.C
		fmt.Println("current time: ", t)
	}
}
