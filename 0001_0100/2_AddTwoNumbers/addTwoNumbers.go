package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// // 解法1
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	result := l1
// 	add := 0
// 	lastPtr := l1
// 	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
// 		l1.Val += l2.Val + add
// 		if l1.Val < 10 {
// 			add = 0
// 		} else {
// 			l1.Val -= 10
// 			add = 1
// 		}
// 		lastPtr = l1
// 	}
// 	if l2 != nil {
// 		lastPtr.Next = l2
// 		l1 = lastPtr.Next
// 	}
// 	for ; l1 != nil && add != 0; l1 = l1.Next {
// 		l1.Val += add
// 		if l1.Val < 10 {
// 			add = 0
// 		} else {
// 			l1.Val -= 10
// 			add = 1
// 		}
// 		lastPtr = l1
// 	}
// 	if add == 1 {
// 		lastPtr.Next = &ListNode{1, nil}
// 	}
// 	return result
// }

// 另一个方法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := &ListNode{0, l1}
	p2 := &ListNode{0, l2}
	forward := 0
	for p1.Next != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
		p1.Val += p2.Val + forward
		if p1.Val < 10 {
			forward = 0
		} else {
			p1.Val -= 10
			forward = 1
		}
	}
	if p2.Next != nil {
		p1.Next = p2.Next
	}
	for p1.Next != nil && forward == 1 {
		p1 = p1.Next
		p1.Val += forward
		if p1.Val < 10 {
			forward = 0
		} else {
			p1.Val -= 10
			forward = 1
		}
	}
	if forward == 1 {
		p1.Next = &ListNode{1, nil}
	}
	return l1
}
