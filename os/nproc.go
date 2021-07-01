package main

import (
	"fmt"
	"runtime"
)

// goroutines are distributed on runtime.NumCPU() CPUs
// The GOMAXPROCS variable limits the number of operating system threads
// that can execute user-level Go code simultaneously
// therefore, GOMAXPROCS can be larger than the actual number of CPUs
func main() {
	fmt.Println(runtime.GOMAXPROCS(-1)) // 4
	fmt.Println(runtime.NumCPU())       // 4
	fmt.Println(runtime.NumGoroutine()) // 1
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(-1)) // 20
	runtime.GOMAXPROCS(300)
	fmt.Println(runtime.GOMAXPROCS(-1)) // 300
	fmt.Printf("runtime: %s\narchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}
