package main

import (
	"fmt"
	"reflect"
)

type ABC struct {
	a int
	b string
	c []int
}

func main() {
	var a = ABC{a: 1, b: "10", c: []int{1, 2}}
	var b = ABC{a: 1, b: "10", c: []int{1, 2}}
	fmt.Println(reflect.DeepEqual(a, b)) // true

	c, d := 1, 1
	fmt.Println(&c == &d)                  // false
	fmt.Println(reflect.DeepEqual(&c, &d)) // true
}
