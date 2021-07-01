package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach")) // true

	fmt.Println(r.FindString("peach punch")) // peach

	fmt.Println(r.FindStringIndex("peach punch")) // [0, 5] ([start, end))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH
}
