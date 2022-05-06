package mergeKLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// // 使用最小堆的解法
// type minHeap [](*ListNode)

// func (h minHeap) Len() int {
// 	return len(h)
// }
// func (h minHeap) Less(i, j int) bool {
// 	return h[i].Val <= h[j].Val
// }
// func (h minHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }
// func (h *minHeap) Push(x interface{}) {
// 	*h = append(*h, x.(*ListNode))
// }
// func (h *minHeap) Pop() interface{} {
// 	stackTopIndex := len(*h) - 1
// 	stackTop := (*h)[stackTopIndex]
// 	(*h) = (*h)[:stackTopIndex]
// 	return stackTop
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	for i := 0; i < len(lists); i++ {
// 		if lists[i] == nil {
// 			lists = append(lists[:i], lists[i+1:]...)
// 		}
// 	}
// 	minheap := minHeap(lists)
// 	h := &minheap
// 	heap.Init(h)
// 	dummyHead := &ListNode{}
// 	temp := dummyHead
// 	for h.Len() != 0 {
// 		min := heap.Pop(h).(*ListNode)
// 		temp.Next = min
// 		temp = temp.Next
// 		if min.Next != nil {
// 			heap.Push(h, min.Next)
// 		}
// 	}
// 	return dummyHead.Next
// }

// 使用mergeSort的解法
func mergeKLists(lists []*ListNode) *ListNode {
	listLen := len(lists)
	if listLen == 0 {
		return nil
	}
	return mergeListInRange(lists, 0, listLen-1)
}

func mergeListInRange(lists []*ListNode, start int, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	mid := start + (end-start)>>1
	head1 := mergeListInRange(lists, start, mid)
	head2 := mergeListInRange(lists, mid+1, end)
	return merge(head1, head2)
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	} else {
		cur.Next = head2
	}
	return dummyHead.Next
}
