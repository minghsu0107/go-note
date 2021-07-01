package main

import (
	"fmt"
	"math"
)

/*
int8 int16 int32 int64
uint8 uint16 uint32 uint64
float32 float64
int, uint: depends on the operating system

type rune int32
type byte uint8
*/

// we can DECLARE a VARIABLE is of a certain TYPE it can only hold VALUES of that TYPE
var z int

// z = 21 // cannot do that

type hotdog int

type myFunc func(a int, b int)

func main() {
	// can only assign a value to a declared the variable within a function
	z = 21
	fmt.Println(z)

	var foo int = 42
	var bar hotdog
	// bar = foo // cannot assign the value of a type int to a type hotdog
	bar = hotdog(foo) // we can convert the value of a to a value of type hotdog

	// however, we can do that to functions
	fun := func(a int, b int) {}
	h1 := myFunc(fun)
	h2 := fun
	h1 = h2
	// data type still the same
	fmt.Printf("%T\n", h1) // main.myFunc
	h2 = myFunc(fun)
	fmt.Printf("%T\n", h2) // func(int, int)

	fmt.Println(foo)
	fmt.Printf("%T\n", foo)
	fmt.Println(bar)
	fmt.Printf("%T\n", bar) // main.hotdog

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) // 0

	// you can only use the short declaration inside of a code block
	// you have to use the variable you create
	f := "apple"
	// var f string = "apple"
	fmt.Println(f)

	const s string = "constant"
	fmt.Println(s)

	const n = 500000000
	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call.
	// For example, here math.Sin expects a float64
	fmt.Println(math.Sin(n))

	const g = 3e20 / n
	fmt.Println(g)        // 6e+11
	fmt.Println(int64(g)) // 600000000000

	var c1 byte = 'a'
	var c2 int = '北'              // avoiding overflow
	fmt.Println(c1)               // 97
	fmt.Printf("%c\n", c1)        // a
	fmt.Printf("%c %d\n", c2, c2) // 北 21271

	c3 := "I love 台灣"
	fmt.Println(c3[2])         //108
	fmt.Println(string(c3[2])) // l
	fmt.Println(string(c3[7])) // 亂碼 三個byte才能形成中文
	fmt.Println([]byte(c3))    // [73 32 108 111 118 101 32 229 143 176 231 129 163]

	c4 := []rune(c3)
	fmt.Println(c4)            // [73 32 108 111 118 101 32 21488 28771]
	fmt.Println(string(c4[7])) // 台
	fmt.Println(string(c4[8])) // 灣

	var dd uint8 = 2
	fmt.Printf("%08b\n", dd)  // 00000010
	fmt.Printf("%08b\n", ^dd) // 11111101
}
