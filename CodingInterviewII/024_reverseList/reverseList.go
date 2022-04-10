package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var preNode, temp *ListNode
	for {
		temp = head.Next
		head.Next = preNode
		preNode = head
		if temp == nil {
			break
		}
		head = temp
	}
	return preNode
}
