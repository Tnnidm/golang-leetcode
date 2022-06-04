package findKthLargest

import "container/heap"

type minHeap []int

func (mh minHeap) Len() int {
	return len(mh)
}
func (mh minHeap) Less(i, j int) bool {
	return mh[i] <= mh[j]
}
func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}
func (mh *minHeap) Pop() interface{} {
	topIndex := mh.Len() - 1
	temp := (*mh)[topIndex]
	(*mh) = (*mh)[:topIndex]
	return temp
}

func findKthLargest(nums []int, k int) int {
	minheap := minHeap(nums[:k])
	mh := &minheap
	heap.Init(mh)
	for i := k; i < len(nums); i++ {
		if nums[i] > (*mh)[0] {
			heap.Pop(mh)
			heap.Push(mh, nums[i])
		}
	}
	return (*mh)[0]
}
