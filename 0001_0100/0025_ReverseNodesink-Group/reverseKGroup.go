package reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	temp := head
	// var preNext *ListNode
	dummyHead := &ListNode{}
	for i := 0; i < k; i++ {
		if temp == nil {
			return head
		}
		next := temp.Next
		dummyHeadNext := dummyHead.Next
		dummyHead.Next = temp
		temp.Next = dummyHeadNext
		// preNext = temp
		temp = next
	}
	head.Next = reverseKGroup(temp, k)
	return dummyHead.Next
}
