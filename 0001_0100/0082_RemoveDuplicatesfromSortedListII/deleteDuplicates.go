package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

// // 递归方法
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	if head.Val != head.Next.Val {
// 		head.Next = deleteDuplicates(head.Next)
// 		return head
// 	} else {
// 		value := head.Val
// 		for head != nil && head.Val == value {
// 			head = head.Next
// 		}
// 		return deleteDuplicates(head)
// 	}
// }

// 迭代方法
func deleteDuplicates(head *ListNode) *ListNode {
	assistHead := &ListNode{101, head}
	for p := assistHead; p.Next != nil && p.Next.Next != nil; {
		if p.Next.Val == p.Next.Next.Val {
			sameValue := p.Next.Val
			for p.Next != nil && p.Next.Val == sameValue {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	return assistHead.Next
}
