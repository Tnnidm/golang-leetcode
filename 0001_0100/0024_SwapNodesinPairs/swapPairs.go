package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// // 迭代方法
// func swapPairs(head *ListNode) *ListNode {
// 	assist := &ListNode{0, head}
// 	for p := assist; p.Next != nil && p.Next.Next != nil; p = p.Next.Next {
// 		behand3 := p.Next.Next.Next
// 		p.Next.Next.Next = p.Next
// 		p.Next = p.Next.Next
// 		p.Next.Next.Next = behand3
// 	}
// 	return assist.Next
// }

// 方法2: 递归方法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := head.Next.Next
	newhead := head.Next
	newhead.Next = head
	head.Next = swapPairs(last)
	return newhead
}
