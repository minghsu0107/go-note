package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

const numWorkers = 3
const numJobs = 10
const batchLines = 3

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jobs := make(chan []string, numJobs)

	wg := sync.WaitGroup{}
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(jobs, &wg)
	}

	scanFile(file, jobs)

	close(jobs)
	wg.Wait()

}

func scanFile(file *os.File, jobs chan []string) {
	scanner := bufio.NewScanner(file)

	batch := make([]string, 0, batchLines)
	for scanner.Scan() {
		line := scanner.Text()
		batch = append(batch, line)

		if len(batch) >= batchLines {
			jobs <- batch
			batch = make([]string, 0, batchLines)
		}
	}

	if len(batch) > 0 {
		jobs <- batch
	}
}

func worker(jobs chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for batch := range jobs {
		// process batch
		printBatch(batch)
	}
}

func printBatch(batch []string) {
	fmt.Printf("Batch size: %d, content: %s\n", len(batch), batch)
}
