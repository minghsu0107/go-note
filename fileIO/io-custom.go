package main

import (
	"fmt"
	"io"
)

// custom reader
// io.Reader is an interface
// we only need to implement Read method
type Ustr struct {
	s string // data
	i int    // offset
}

func NewUstr(s string) *Ustr {
	return &Ustr{s, 0}
}

// length of the unread string
func (s *Ustr) Len() int {
	return len(s.s) - s.i
}

func (s *Ustr) Read(p []byte) (n int, err error) {
	for ; s.i < len(s.s) && n < len(p); s.i++ {
		c := s.s[s.i]
		// transform lower case to upper case
		if 'a' <= c && c <= 'z' {
			p[n] = 'A' + c - 'a'
		} else {
			p[n] = c
		}
		n++
	}
	if n == 0 {
		return n, io.EOF
	}
	return n, nil
}
func main() {
	s := NewUstr("Hello World!")
	buf := make([]byte, s.Len())

	if n, err := io.ReadFull(s, buf); err == io.EOF {
		fmt.Println("len(s) = 0")
	} else {
		fmt.Printf("%s\n", buf) // HELLO WORLD!
		fmt.Println(n, err)     // 12 <nil>
	}
}
