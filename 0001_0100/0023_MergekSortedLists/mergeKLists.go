package mergeKLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// // 最小堆方法
// type MinHeap [](*ListNode)

// func (mh MinHeap) Len() int {
// 	return len(mh)
// }

// func (mh MinHeap) Less(i, j int) bool {
// 	return mh[i].Val < mh[j].Val
// }

// func (mh MinHeap) Swap(i, j int) {
// 	mh[i], mh[j] = mh[j], mh[i]
// }

// func (mh *MinHeap) Push(x interface{}) {
// 	*mh = append(*mh, x.(*ListNode))
// }

// func (mh *MinHeap) Pop() interface{} {
// 	lastIndex := mh.Len() - 1
// 	temp := (*mh)[lastIndex]
// 	*mh = (*mh)[:lastIndex]
// 	return temp
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	listswoNil := make([](*ListNode), 0, len(lists))
// 	for _, list := range lists {
// 		if list != nil {
// 			listswoNil = append(listswoNil, list)
// 		}
// 	}
// 	minHeap := MinHeap(listswoNil)
// 	mh := &minHeap
// 	heap.Init(mh)
// 	dummyHead := &ListNode{}
// 	temp := dummyHead
// 	for mh.Len() != 0 {
// 		minimal := heap.Pop(mh).(*ListNode)
// 		temp.Next = minimal
// 		temp = temp.Next
// 		if minimal.Next != nil {
// 			heap.Push(mh, minimal.Next)
// 		}
// 	}
// 	return dummyHead.Next
// }

// 分治方法
func mergeKLists(lists []*ListNode) *ListNode {
	listNum := len(lists)
	if listNum == 0 {
		return nil
	} else if listNum == 1 {
		return lists[0]
	} else {
		return mergeCore(mergeKLists(lists[:listNum>>1]), mergeKLists(lists[listNum>>1:]))
	}
}

func mergeCore(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp := dummyHead
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			temp.Next = head1
			head1 = head1.Next
		} else {
			temp.Next = head2
			head2 = head2.Next
		}
		temp = temp.Next
	}
	if head1 != nil {
		temp.Next = head1
	} else {
		temp.Next = head2
	}
	return dummyHead.Next
}
