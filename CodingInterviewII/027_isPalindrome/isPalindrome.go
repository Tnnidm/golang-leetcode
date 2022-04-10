package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l2 := slow.Next
	slow.Next = nil
	l2 = reverseList(l2)
	for l2 != nil {
		if head.Val != l2.Val {
			return false
		}
		head = head.Next
		l2 = l2.Next
	}
	return true
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
