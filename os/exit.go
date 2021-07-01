package main

import (
	"fmt"
	"os"
)

func main() {
	// defers will not be run when using os.Exit
	defer fmt.Println("this will not be printed")

	os.Exit(3)
}
