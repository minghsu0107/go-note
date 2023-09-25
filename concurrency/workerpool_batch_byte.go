package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

const numWorkers = 3
const numJobs = 10
const readByteSize = 10
const batchByteSize = 100

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jobs := make(chan []byte, numJobs)

	wg := sync.WaitGroup{}
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(jobs, &wg)
	}

	batchFile(file, jobs)

	close(jobs)
	wg.Wait()
}

func batchFile(file *os.File, jobs chan []byte) {
	batch := make([]byte, 0, batchByteSize)
	b := make([]byte, readByteSize)

	for {
		bytesRead, err := file.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		batch = append(batch, b[:bytesRead]...)
		if len(batch) >= batchByteSize {
			jobs <- batch
			batch = make([]byte, 0, batchByteSize) // new empty batch buffer
		}
	}

	if len(batch) > 0 {
		jobs <- batch
	}

}

func worker(jobs chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range jobs {
		fmt.Printf("Read %d bytes, content: %s\n", len(data), string(data))
	}
}
