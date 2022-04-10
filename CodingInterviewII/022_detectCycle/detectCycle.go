package detectCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for fast != nil {
				if slow == fast {
					return slow
				}
				fast = fast.Next
				slow = slow.Next
			}
		}
	}
	return nil
}
