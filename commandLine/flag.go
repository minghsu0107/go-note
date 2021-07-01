package main

import (
	"flag"
	"fmt"
)

// Use -h or --help flags to get automatically generated help text
func main() {

	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// It’s also possible to declare an option
	// that uses an existing var declared elsewhere in the program
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	// Trailing positional arguments can be provided after any flags
	// otherwise the flags will be interpreted as positional arguments
	fmt.Println("tail:", flag.Args())
}

/*
$ ./flag -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

$ ./flag -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

$ ./flag -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

$ ./flag -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string
*/
