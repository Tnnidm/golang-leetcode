package partition

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummyHead := &ListNode{}
	temp1 := dummyHead
	dummyHead2 := &ListNode{0, head}
	temp2 := dummyHead2
	for temp2.Next != nil {
		if temp2.Next.Val < x {
			temp1.Next = temp2.Next
			temp1 = temp1.Next
			temp2.Next = temp2.Next.Next
		} else {
			temp2 = temp2.Next
		}
	}
	temp1.Next = dummyHead2.Next
	return dummyHead.Next
}
