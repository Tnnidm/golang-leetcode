package deleteNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, target *ListNode) *ListNode {
	if head == nil || target == nil {
		return head
	}
	if head == target {
		return head.Next
	}
	if target.Next != nil {
		target.Val = target.Next.Val
		target.Next = target.Next.Next
	} else {
		p := head
		for p.Next != target {
			p = p.Next
		}
		p.Next = nil
	}
	return head
}
