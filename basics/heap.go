package main

import (
	"container/heap"
	"fmt"
)

// IntSlice 實現了 heap.Interface
type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *IntSlice) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *IntSlice) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

// IntHeap 封裝了 IntSlice
type IntHeap struct {
	elems IntSlice
}

// 實現相關函式或方法時，透過 heap 提供的函式
func NewIntHeap(numbers ...int) *IntHeap {
	h := &IntHeap{IntSlice(numbers)}
	heap.Init(&(h.elems))
	return h
}

func (h *IntHeap) Push(n int) {
	heap.Push(&(h.elems), n)
}

func (h *IntHeap) Pop() int {
	return heap.Pop(&(h.elems)).(int)
}

func (h *IntHeap) Len() int {
	return len(h.elems)
}

// 一律透過 h 來操作
func main() {
	h := NewIntHeap(2, 1, 5)
	h.Push(3)
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
}
