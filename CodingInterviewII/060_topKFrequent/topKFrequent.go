package topKFrequent

import "container/heap"

type maxHeap struct {
	kmh  *([]int)
	freq map[int]int
}

func (h maxHeap) Len() int {
	return len(*h.kmh)
}

func (h maxHeap) Less(i, j int) bool {
	return h.freq[(*h.kmh)[i]] < h.freq[(*h.kmh)[j]]
}

func (h maxHeap) Swap(i, j int) {
	(*h.kmh)[i], (*h.kmh)[j] = (*h.kmh)[j], (*h.kmh)[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h.kmh = append(*h.kmh, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	lastIndex := len(*h.kmh) - 1
	top := (*h.kmh)[lastIndex]
	*h.kmh = (*h.kmh)[:lastIndex]
	return top
}

func topKFrequent(nums []int, k int) []int {
	mh := &maxHeap{kmh: &[]int{}, freq: map[int]int{}}
	for _, num := range nums {
		(*mh).freq[num] += 1
	}
	for num := range (*mh).freq {
		heap.Push(mh, num)
		if (*mh).Len() > k {
			heap.Pop(mh)
		}
	}
	return *mh.kmh
}
