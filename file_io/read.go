package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(reflect.TypeOf(dat)) // []uint8
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat") // f is *File (pointer)
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	// n3, err := io.ReadFull(f, b3)
	n3, err := io.ReadAtLeast(f, b3, 2) // robust
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	// The bufio package implements a buffered reader
	// that may be useful both for its efficiency with many small reads
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	b5, err := r4.ReadString('\n') // read a string (until delim: \n)
	f.Close()
}
