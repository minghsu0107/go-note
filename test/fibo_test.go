package main

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	t.Skip()
}

func TestFibo(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{3, 2},
		{4, 3},
		{5, 5},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("n=%d", tt.n)
		t.Run(testname, func(t *testing.T) {
			ans := Fibo(tt.n)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibo(5)
	}
}
