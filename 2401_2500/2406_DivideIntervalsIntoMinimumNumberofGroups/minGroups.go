package minGroups

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *minHeap) Pop() interface{} {
	lastIndex := (*mh).Len() - 1
	temp := (*mh)[lastIndex]
	*mh = (*mh)[:lastIndex]
	return temp
}

func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] == intervals[j][1])
	})
	mh := &minHeap{}
	heap.Push(mh, intervals[0][1])
	for i := 1; i < len(intervals); i++ {
		if (*mh)[0] < intervals[i][0] {
			heap.Pop(mh)
			heap.Push(mh, intervals[i][1])
		} else {
			heap.Push(mh, intervals[i][1])
		}
	}
	return mh.Len()
}
