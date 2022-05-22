package mergeKLists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap [](*ListNode)

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Val < mh[j].Val
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(*ListNode))
}

func (mh *MinHeap) Pop() interface{} {
	lastIndex := mh.Len() - 1
	temp := (*mh)[lastIndex]
	*mh = (*mh)[:lastIndex]
	return temp
}

func mergeKLists(lists []*ListNode) *ListNode {
	listswoNil := make([](*ListNode), 0, len(lists))
	for _, list := range lists {
		if list != nil {
			listswoNil = append(listswoNil, list)
		}
	}
	minHeap := MinHeap(listswoNil)
	mh := &minHeap
	heap.Init(mh)
	dummyHead := &ListNode{}
	temp := dummyHead
	for mh.Len() != 0 {
		minimal := heap.Pop(mh).(*ListNode)
		temp.Next = minimal
		temp = temp.Next
		if minimal.Next != nil {
			heap.Push(mh, minimal.Next)
		}
	}
	return dummyHead.Next
}
