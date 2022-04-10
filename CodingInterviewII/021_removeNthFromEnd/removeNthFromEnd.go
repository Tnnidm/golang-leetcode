package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	assistHead := &ListNode{0, head}
	slow, fast := assistHead, assistHead
	for n != 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return assistHead.Next
}
