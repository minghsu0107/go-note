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
	jobs := make(chan int, numJobs)
	results := make(chan int)

	var wg sync.WaitGroup
	for w := 1; w <= workerPoolSize; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	cnt := 0
	for res := range results {
		fmt.Println(res)
		cnt++
		// if we know the length of total results, we can use unbuffered channel
		if cnt == 5 {
			break
		}
	}
	close(results)
	wg.Wait()
}
