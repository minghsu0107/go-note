package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HttpDoTest(ctx context.Context, resChan chan<- string) error {
	start := time.Now()

	repoUrl := "https://api.github.com/repos/campoy/golang-plugins"
	req, err := http.NewRequest("GET", repoUrl, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest Error: %s", err.Error())
	}

	// in go >= 1.7
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do Error: %s", err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll Error: %s", err.Error())
	}

	log.Printf("Read body size [%d]", len(data))
	log.Println("CostTime is: " + time.Since(start).String())

	resChan <- string(data)

	return nil
}

func main() {
	deadline := 1.5
	d := time.Now().Add(time.Duration(deadline) * time.Second) // deadline max
	// Background returns a non-nil, empty Context (for initialization)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// when time is up, cancel() will be called
	// done channel is closed when the returned cancel function is called
	// or when the parent context's Done channel is closed
	// this ensures that the done channel is always closed eventually
	defer cancel()

	resChan := make(chan string)

	go HttpDoTest(ctx, resChan)

	var resData string
	select {
	// will return when done channel (ctx.Done()) is closed
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case resData = <-resChan:
		fmt.Println("Read data finished")
	}

	log.Printf("Read data size: [%d]", len(resData))
}

/*
context deadline exceeded
2020/07/16 01:11:10 Read data size: [0]
*/
