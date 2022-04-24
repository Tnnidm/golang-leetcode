package heap

import "container/heap"

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	lastIndex := len(*h) - 1
	lastOne := (*h)[lastIndex]
	(*h) = (*h)[:lastIndex]
	return lastOne
}

func usage() {
	h := &maxHeap{}
	heap.Push(h, 1)
	heap.Pop(h)
}
