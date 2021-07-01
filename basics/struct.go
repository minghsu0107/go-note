package main

import "fmt"

// only fields starting with a capital letter are exported (visible) outside the curent package
// so outside the "main" package, person.name is invalid
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	fmt.Println(newPerson("Jon")) // &{Jon 42}

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // 50

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age) // 51

	type data struct {
		name string
	}
	a := []data{{"Tom"}}
	a[0].name = "Jerry"
	fmt.Println(a) // [{Jerry}]

	// to change the struct in a map, we should use pointer
	m := map[string]*data{
		"x": {"Tom"},
	}
	m["x"].name = "Jerry"
	fmt.Println(m["x"]) // &[{Jerry}]
}
