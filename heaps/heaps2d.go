package main

type data struct {
	time  int
	count int
}

type MaxHeap []data

// Heap Interface : Sort.interface

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i int, j int) bool {
	return h[i].time > h[j].time
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(data))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1:n]
	return x
}
