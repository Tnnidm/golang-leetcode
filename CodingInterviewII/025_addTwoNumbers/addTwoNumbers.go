package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	len1, len2 := 0, 0
	for temp := l1; temp != nil; temp = temp.Next {
		len1++
	}
	for temp := l2; temp != nil; temp = temp.Next {
		len2++
	}
	if len1 < len2 {
		l1, l2 = l2, l1
	}
	addOne := 0
	temp := l1
	for {
		if l2 != nil {
			temp.Val += l2.Val
			l2 = l2.Next
		}
		temp.Val += addOne
		if temp.Val >= 10 {
			temp.Val -= 10
			addOne = 1
		} else {
			addOne = 0
		}
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	if addOne == 1 {
		temp.Next = &ListNode{1, nil}
	}
	return reverseList(l1)
}

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
