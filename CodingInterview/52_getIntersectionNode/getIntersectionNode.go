package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptr1 := headA
	ptr2 := headB
	for ptr1 != nil || ptr2 != nil {
		if ptr1 == ptr2 {
			return ptr1
		}
		if ptr1 == nil {
			ptr1 = headB
		} else {
			ptr1 = ptr1.Next
		}
		if ptr2 == nil {
			ptr2 = headA
		} else {
			ptr2 = ptr2.Next
		}
	}
	return nil
}
