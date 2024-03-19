package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func main() {

	hp := &MinHeap{1, 2, 4, 5, 6, 4, 3, 1, 7, 8}
	heap.Init(hp)
	fmt.Println(*hp)

	hp2 := &MaxHeap{data{time: 20, count: 40}, data{time: 10, count: 40}, data{time: 60, count: 40}, data{time: 30, count: 40}}

	heap.Init(hp2)
	fmt.Println(*hp2)
	for i := 0; i < 4; i++ {
		fmt.Println(hp2.Pop().(data))
	}

}

// Heap Interface : sort.Interface

// Len()
func (h MinHeap) Len() int {
	return len(h)
}

// Swap

func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// Less : use this function to change the heap to MinHeap/MaxHeap

func (h MinHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

// Heap Interface : Push

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Heap Interface : Pop()
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1 : n-1]
	return x
}
