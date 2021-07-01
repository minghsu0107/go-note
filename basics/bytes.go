package main

import (
	"bytes"
	"fmt"
)

func main() {

	b1 := []byte("hello world!")
	buf := bytes.NewBuffer(b1) // buf is a pointer (*Buffer)
	fmt.Printf("buff len=%d\n", buf.Len()) // 12
	fmt.Printf("buff cap=%d\n", buf.Cap()) // 12

	buf.Grow(100)
	fmt.Printf("buff len=%d\n", buf.Len()) // 12
	fmt.Printf("buff cap=%d\n", buf.Cap()) // 124

	b2 := make([]byte, 6)
	buf.Read(b2)        // read data in buf to b2
	println(string(b2)) //hello

	b3 := buf.Next(5)
	println(string(b3)) //world

	b4 := buf.Next(3)
	println(string(b4)) // !

	buf2 := bytes.NewBuffer(b1)
	// ReadBytes reads until the first occurrence of delim in the input,
	// returning a slice containing the data up to and including the delimiter.
	b5, _ := buf2.ReadBytes(byte(' '))
	println(len(b5))    // 6
	println(string(b5)) // hello

	b6 := []byte("go programming")
	buf3 := bytes.NewBuffer(b1)
	// Write appends the contents of b6 to the buffer, growing the buffer as
	// needed.
	buf3.Write(b6)
	println(string(buf3.Bytes())) // hello world!go programming
}
