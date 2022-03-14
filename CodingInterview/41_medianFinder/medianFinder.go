package MedianFinder

import "container/heap"

// 方法1：最大堆最小堆
type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}
func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *maxHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}
func (m *maxHeap) Pop() interface{} {
	mLen := len(*m)
	temp := (*m)[mLen-1]
	(*m) = (*m)[:mLen-1]
	return temp
}

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *minHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}
func (m *minHeap) Pop() interface{} {
	mLen := len(*m)
	temp := (*m)[mLen-1]
	(*m) = (*m)[:mLen-1]
	return temp
}

type MedianFinder struct {
	maxH *maxHeap
	minH *minHeap
}

func Constructor() MedianFinder {
	return MedianFinder{&maxHeap{}, &minHeap{}}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.maxH.Len() == mf.minH.Len() {
		if mf.maxH.Len() == 0 || num <= (*mf.maxH)[0] {
			heap.Push(mf.maxH, num)
		} else {
			heap.Push(mf.minH, num)
			top := heap.Pop(mf.minH).(int)
			heap.Push(mf.maxH, top)
		}
	} else {
		if num > (*mf.maxH)[0] {
			heap.Push(mf.minH, num)
		} else {
			heap.Push(mf.maxH, num)
			top := heap.Pop(mf.maxH).(int)
			heap.Push(mf.minH, top)
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxH.Len() == mf.minH.Len() {
		return float64((*mf.maxH)[0]+(*mf.minH)[0]) / 2
	} else {
		return float64((*mf.maxH)[0])
	}
}
