package rotateRight

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	count := 1
	p := head
	for ; p.Next != nil; p = p.Next {
		count++
	}
	p.Next = head
	k = k % count
	for i := 0; i < count-k; i++ {
		p = p.Next
	}
	temp := p.Next
	p.Next = nil
	return temp
}
