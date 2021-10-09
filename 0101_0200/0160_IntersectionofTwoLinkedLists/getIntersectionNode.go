package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

// // HashMap方法
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	aMap := map[*ListNode]struct{}{}
// 	for headA != nil {
// 		aMap[headA] = struct{}{}
// 		headA = headA.Next
// 	}
// 	for headB != nil {
// 		if _, ok := aMap[headB]; ok {
// 			return headB
// 		}
// 		headB = headB.Next
// 	}
// 	return nil
// }

// 双指针方法，很巧妙的一种方法
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
