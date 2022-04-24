package kthLargest

import "container/heap"

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	lastIndex := len(*h) - 1
	lastOne := (*h)[lastIndex]
	*h = (*h)[:lastIndex]
	return lastOne
}

type KthLargest struct {
	kmh *minHeap
	k   int
}

//
func Constructor(k int, nums []int) KthLargest {
	h := minHeap(nums)
	heap.Init(&h)
	for i := len(nums); i > k; i-- {
		heap.Pop(&h)
	}
	return KthLargest{kmh: &h, k: k}
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl.kmh, val)
	if (*kl.kmh).Len() > kl.k {
		heap.Pop(kl.kmh)
	}
	return (*kl.kmh)[0]
}
