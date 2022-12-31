package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// The program only takes about 2 seconds despite doing about 5 seconds of total work
// because there are 3 workers operating concurrently

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked because there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs and then close that channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// if we don't close the job, the worker will wait for new jobs
	// after finish the current task (does not exit)
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have finished.
	// however, if any worker fails to write result, our main thread will block forever
	// since there won't be numJobs successfully done
	// to solve this, use waitgroup (see worker-pool2.go)
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

/*
worker 1 started  job 2
worker 3 started  job 1
worker 2 started  job 3
worker 2 finished job 3
worker 2 started  job 4
worker 3 finished job 1
worker 3 started  job 5
worker 1 finished job 2
worker 2 finished job 4
worker 3 finished job 5
*/
