package sortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	var preSlow *ListNode
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	preSlow.Next = nil
	head = sortList(head)
	secondPart := sortList(slow)
	dummyHead := &ListNode{}
	temp := dummyHead
	for head != nil && secondPart != nil {
		if head.Val < secondPart.Val {
			temp.Next = head
			temp = temp.Next
			head = head.Next
		} else {
			temp.Next = secondPart
			temp = temp.Next
			secondPart = secondPart.Next
		}
	}
	if head != nil {
		temp.Next = head
	} else {
		temp.Next = secondPart
	}
	return dummyHead.Next
}
