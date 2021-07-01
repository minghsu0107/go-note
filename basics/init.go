package main

import (
	"fmt"
)

// init function can be declared multiple times
// each init function will be called in orders right before main()
func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("Hello, playground")
}

/*
init 1
init 2
Hello, playground
*/

/*
// just import init() of the package
import(
    // Needed for the Postgresql driver
    _ "github.com/lib/pq
)
*/
