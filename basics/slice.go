package main

import "fmt"

/*
Unlike array, slices are typed only by the elements they contain (not the number of elements)
*/
func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s) // initially zero-valued

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slicing still references the original memory location
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	str := []byte("hello")
	fmt.Println(string(str)) // hello

	var lettersLower = []rune("abcdefghijklmnopqrstuvwxyz")
	var lettersUpper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	x := append(lettersLower, lettersUpper...)
	fmt.Println(string(lettersLower)) // does not change
	fmt.Println(string(x))
	fmt.Printf("%p\n", &lettersLower)
	fmt.Printf("%p\n", &x)         // new memory allocation
	fmt.Println(cap(lettersLower)) // 26
	fmt.Println(cap(x))            // 52

	x = append(x, 'a')
	fmt.Println(string(x))
	fmt.Printf("%p\n", &x) // same memory allocation
}
