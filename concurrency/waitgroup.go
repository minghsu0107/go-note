package main

import (
	"fmt"
	"sync"
	"time"
)

// Note that a WaitGroup must be passed to functions by pointer
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done() // On return, notify the WaitGroup that we’re done

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they’re done.
	wg.Wait()
}

/*
Worker 3 starting
Worker 4 starting
Worker 2 starting
Worker 5 starting
Worker 1 starting
Worker 5 done
Worker 2 done
Worker 1 done
Worker 4 done
Worker 3 done
*/
