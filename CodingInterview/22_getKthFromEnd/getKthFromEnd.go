package getKthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for ; k > 0; k-- {
		if fast != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
