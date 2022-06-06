package topKFrequent

import "container/heap"

type minHeap [][2]int

func (mh minHeap) Len() int {
	return len(mh)
}
func (mh minHeap) Less(i, j int) bool {
	return mh[i][1] < mh[j][1]
}
func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.([2]int))
}
func (mh *minHeap) Pop() interface{} {
	lastIndex := mh.Len() - 1
	temp := (*mh)[lastIndex]
	*mh = (*mh)[:lastIndex]
	return temp
}
func topKFrequent(nums []int, k int) []int {
	numCount := make(map[int]int, k)
	for _, num := range nums {
		numCount[num]++
	}
	mh := &minHeap{}
	for num, count := range numCount {
		heap.Push(mh, [2]int{num, count})
		if mh.Len() > k {
			heap.Pop(mh)
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = (*mh)[i][0]
	}
	return result
}
