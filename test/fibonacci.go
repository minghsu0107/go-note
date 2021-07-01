package main

func Fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibo(n-1) + Fibo(n-2)
}
