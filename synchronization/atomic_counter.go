package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
		// Reading atomics safely while they are being updated is also possible
		go func() {
			fmt.Println(atomic.LoadUint64(&ops))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
