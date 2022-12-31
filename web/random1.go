package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 81,87
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100)) // integer in [0, 100)
	fmt.Println()

	fmt.Println(rand.Float64()) // float64 in [0, 1)

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// new seed
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
