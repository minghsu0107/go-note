package main

import (
	"fmt"
	"os"
)

func main() {
	foo := []byte("this is a BIG OLD TEST!!\n")
	tst := []byte("little test")
	bar := append(foo[:10], tst...)

	// now bar is right, but foo is a mix of old and new text!
	fmt.Print("without copy, foo after:  ")
	os.Stdout.Write(foo)

	// ok, now the same exercise but with an explicit copy of foo
	foo = []byte("this is a BIG OLD TEST!!\n")
	bar = append([]byte(nil), foo[:10]...) // copies foo[:10]
	bar = append(bar, tst...)

	// this time we modified a copy, and foo is its original self
	fmt.Print("with a copy, foo after:   ")
	os.Stdout.Write(foo)
}
