package kSmallestPairs

import (
	"container/heap"
)

type minHeap struct {
	heap  [][2]int
	nums1 *[]int
	nums2 *[]int
}

func (h minHeap) Len() int {
	return len(h.heap)
}

func (h minHeap) Less(i, j int) bool {
	return (*h.nums1)[h.heap[i][0]]+(*h.nums2)[h.heap[i][1]] < (*h.nums1)[h.heap[j][0]]+(*h.nums2)[h.heap[j][1]]
}

func (h minHeap) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *minHeap) Push(x interface{}) {
	(*h).heap = append((*h).heap, x.([2]int))
}

func (h *minHeap) Pop() interface{} {
	lastIndex := len((*h).heap) - 1
	lastOne := (*h).heap[lastIndex]
	(*h).heap = (*h).heap[:lastIndex]
	return lastOne
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	nums1Len, nums2Len := len(nums1), len(nums2)
	h := make([][2]int, nums2Len)
	for j := 0; j < nums2Len; j++ {
		h[j] = [2]int{0, j}
	}
	mh := &minHeap{
		heap:  h,
		nums1: &nums1,
		nums2: &nums2,
	}
	heap.Init(mh)
	result := make([][]int, k)
	i := 0
	for ; i < k && mh.Len() != 0; i++ {
		top := heap.Pop(mh).([2]int)
		result[i] = []int{nums1[top[0]], nums2[top[1]]}
		if top[0] < nums1Len-1 {
			heap.Push(mh, [2]int{top[0] + 1, top[1]})
		}
	}
	return result[:i]
}
