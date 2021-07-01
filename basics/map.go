package main

import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Println(m["s"]) // 0

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) // map: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m)) // 2

	delete(m, "k2")
	fmt.Println("map:", m)
	fmt.Println(m["k2"]) // 0

	/*
	   The optional second return value when getting a value from a map
	   indicates if the key was present in the map.

	   This can be used to disambiguate between missing keys
	   and keys with zero values like 0 or ""
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
