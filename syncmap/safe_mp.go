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
			_, loaded := table.LoadOrStore("KEY", n)
			if !loaded {
				fmt.Println("Set to", n) // printed only once
			}
		}(i)
	}
	wg.Wait()
	val, ok := table.Load("KEY")
	fmt.Println(val, ok)
}
