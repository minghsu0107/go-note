package main

import "fmt"

func main() {
	var i = 1
	defer fmt.Println("result: ", func() int { return i * 2 }())
	defer fmt.Println("result: ", func() int { return i * 3 }())
	defer fmt.Println("result: ", func() int { return i * 4 }())
	i++
}

/*
result:  4
result:  3
result:  2
*/
