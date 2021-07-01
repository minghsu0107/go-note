package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g) // print the given object (rect or circle)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
		// should do typw assertion before further operations
		str := val.(string)
		fmt.Println("name: " + str)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(&c)

	// we cannot directly pass names to Printall(), since it accepts only "interface" type
	// an interface value is actually constructed of two words of data;
	// one word is used to point to a method table for the value’s underlying type,
	// and the other word is used to point to the actual data being held by that value
	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
	// vals = append(vals, names...) // invalid
	for i, v := range names {
		vals[i] = v
	}
	vals = append(vals, "ming")
	PrintAll(vals)

	// type assertion
	var i interface{} = "hello"
	if s, ok := i.(string); ok {
		fmt.Println(s + " world") // hello world
	}
	args := []interface{}{
		1,
		"1",
		"2",
	}
	for _, arg := range args {
		if val, ok := arg.(int); ok {
			fmt.Println(val)
		} else {
			fmt.Println("not int.")
		}
	}

	var t interface{}
	t = 3
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

/*
real world usage

type Entity interface {
    UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, v Entity) error {
    return v.UnmarshalHTTP(r)
}

func (u *User) UnmarshalHTTP(r *http.Request) error {
   // ...
}

// usage
var u User
if err := GetEntity(req, &u); err != nil {
    // ...
}
*/
