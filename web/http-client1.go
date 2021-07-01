package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	// can use get moethod to download files
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var status string = resp.Status
	fmt.Println("Response status:", status)
	var header map[string][]string = resp.Header
	// k, v are both string
	for k, v := range header {
		fmt.Print(k, ": ")
		fmt.Println(v[0])
	}

	// note that headers of the body, if any, are with the keys canonicalized
	fmt.Println(resp.Header.Get("Content-Type"))

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(bodyBytes)) // []uint8
	bodyString := string(bodyBytes)
	fmt.Print(bodyString)

	/*
	   // Print the first 5 lines of the response body
	   scanner := bufio.NewScanner(resp.Body)
	   for i := 0; scanner.Scan() && i < 5; i++ {
	       fmt.Println(scanner.Text())
	   }

	   if err := scanner.Err(); err != nil {
	       panic(err)
	   }
	*/
}
