package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var a int
var b string = "James Bond"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Printf("%v\n", a)  // 0
	fmt.Printf("%v\n", b)  // James Bond
	fmt.Printf("%#v\n", a) // 0
	fmt.Printf("%#v\n", b) // "James Bond"
	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", b)  // String
	fmt.Printf("%T\t%T\n", a, b)

	s := fmt.Sprint(a, " something more ", b) // 0 something more James Bond
	fmt.Println(s)
	s2 := fmt.Sprintf("%v\t%T\t%T\n", "to pass in", a, b) // to pass in	   int	 string
	fmt.Println(s2)

	tst := "string"
	tst2 := 10
	tst3 := 1.2

	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))

	// int to string
	var i rune = 100 // rune is the alias of int32
	var j int = 100
	str1 := fmt.Sprintf("%d", i) // %d can be int, int32, int64
	fmt.Println(str1)
	str2 := strconv.Itoa(j) // can only be int
	fmt.Println(str2)
	str3 := strconv.FormatInt(int64(i), 10) // i can be int, int32, int64
	fmt.Println(str3)

	// string to int
	myInt64, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Println(reflect.TypeOf(myInt64)) // int64
	tmp, _ := strconv.ParseInt(str1, 10, 32)
	myInt32 := int32(tmp)
	fmt.Println(reflect.TypeOf(myInt32)) // int32
	myInt, err := strconv.Atoi(str1)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(myInt)) // int

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(m1, m2)) // true

	a1 := []int{1, 2, 3}
	a2 := []int{1, 2, 3}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(a1, a2)) // true

	var mystr = "one"
	var in interface{} = "one"
	fmt.Println("str == in: ", reflect.DeepEqual(mystr, in)) // true

	person1 := person{}
	person2 := person{}
	fmt.Println("person1 == person2: ", reflect.DeepEqual(person1, person2)) // true
}
