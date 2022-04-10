package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tempA, tempB := headA, headB
	for tempA != tempB {
		if tempA == nil {
			tempA = headB
		} else {
			tempA = tempA.Next
		}
		if tempB == nil {
			tempB = headA
		} else {
			tempB = tempB.Next
		}
	}
	return tempA
}
