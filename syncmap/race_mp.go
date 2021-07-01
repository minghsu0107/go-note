package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var table sync.Map

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if _, ok := table.Load("KEY"); !ok {
				table.Store("KEY", n)
				fmt.Println("Set to", n) // be printed more than once
			}
		}(i)
	}
	wg.Wait()
	val, ok := table.Load("KEY")
	fmt.Println(val, ok)
}
