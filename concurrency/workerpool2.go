package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const workerPoolSize = 3
	const numJobs = 5
	jobs := make(chan int, numJobs)    // can be buffered or unbuffered since there are always workers listening to the channel
	results := make(chan int, numJobs) // should be buffered, otherwise will cause dead lock when we close results channel

	var wg sync.WaitGroup
	for w := 1; w <= workerPoolSize; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()

	// close the results channel to indicate that no more values will be sent on it
	// otherwise the main thread will block, waiting for new data sent into results channel
	close(results)
	for res := range results {
		fmt.Println(res)
	}
}
