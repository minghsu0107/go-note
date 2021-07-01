package main

import "fmt"

type rect struct {
	width, height int
}

// Methods can be defined for either pointer or value receiver types

// You may want to use a pointer receiver type
// to avoid copying on method calls
// or to allow the method to mutate the receiving struct
// note that a pointer type can access the methods of its associated value type,
// but not vice versa (*rect can use methods defined on rect, but not vice versa)
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	// Go automatically handles conversion between values and pointers for method calls
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())  // 50
	fmt.Println("perim:", r.perim()) // 30

	rp := &r
	fmt.Println("area: ", rp.area())  // 50
	fmt.Println("perim:", rp.perim()) // 30
}
