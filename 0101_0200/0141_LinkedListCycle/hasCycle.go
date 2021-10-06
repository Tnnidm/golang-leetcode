package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// // HashMap
// func hasCycle(head *ListNode) bool {
// 	nodeMap := map[*ListNode]struct{}{}
// 	for head != nil {
// 		if _, ok := nodeMap[head]; ok {
// 			return true
// 		}
// 		nodeMap[head] = struct{}{}
// 		head = head.Next
// 	}
// 	return false
// }

// Fast-Slow Pointers
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// // Change the linked list
// func hasCycle(head *ListNode) bool {
// 	if head == nil || head.Next == nil {
// 		return false
// 	} else {
// 		ptr := head
// 		ptrNext := head.Next
// 		for ptrNext != nil {
// 			ptr.Next = head
// 			ptr = ptrNext
// 			ptrNext = ptrNext.Next
// 			if ptrNext.Next == head {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// }
