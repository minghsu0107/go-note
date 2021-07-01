package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // load ".env" file by default
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// can also use godotenv.Load("myname.env") to load custom name env file

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR")) // BAR:

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], pair[1])
	}
}
