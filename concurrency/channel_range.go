package main

import (
	"fmt"
	"sync"
)

type Person struct {
	ID int
}

func main() {

	queue := make(chan Person, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for person := range queue {
			wg.Add(1)
			// the inner goroutine (inside for-range closure) holds a reference of person object
			// meaning it's possible that multiple goroutine reference to the same person
			// therefore we should copy the data of each person and pass into the goroutine
			go func(person Person) {
				defer wg.Done()
				fmt.Println(person.ID)
			}(person)
		}
		fmt.Println("done")
	}()

	for i := 0; i < 10; i++ {
		queue <- Person{i}
	}
	close(queue) // close the queue so that the worker will not wait for new jobs forever
	wg.Wait()
}
