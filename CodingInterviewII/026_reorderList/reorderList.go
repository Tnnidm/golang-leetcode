package reorderList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l2 := slow.Next
	slow.Next = nil
	l2 = reverseList(l2)
	for l1 := head; l2 != nil; {
		temp1 := l1.Next
		temp2 := l2.Next
		l1.Next = l2
		l2.Next = temp1
		l1 = temp1
		l2 = temp2
	}
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
