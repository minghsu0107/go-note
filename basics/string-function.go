package main

import (
	b "bytes"
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Join(byte):", string(b.Join([][]byte{[]byte("a"), []byte("b")}, []byte("-"))))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1)) // f00
	p("Replace:   ", s.Replace("foo", "o", "0", 1))  // f0o
	p("Split:     ", s.Split("a-b-c-d-e", "-"))      // [a b c d e]
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
