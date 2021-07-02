package main

import "fmt"

func main() {
  jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for j := range jobs {
			fmt.Println("received job", j)
		}
		done <- true
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
