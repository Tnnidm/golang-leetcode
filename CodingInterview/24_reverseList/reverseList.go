package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var lastNode, nextNode *ListNode
	for head != nil {
		nextNode = head.Next
		head.Next = lastNode
		lastNode = head
		head = nextNode
	}
	return lastNode
}
