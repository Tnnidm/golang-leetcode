package reversePrint

type ListNode struct {
	Val  int
	Next *ListNode
}

// // 实现1
// func reversePrint(head *ListNode) []int {
// 	ans := []int{}
// 	for head != nil {
// 		ans = append(ans, head.Val)
// 		head = head.Next
// 	}
// 	ansLen := len(ans)
// 	if ansLen >= 2 {
// 		for i := 0; i < ansLen/2; i++ {
// 			ans[i], ans[ansLen-1-i] = ans[ansLen-1-i], ans[i]
// 		}
// 	}
// 	return ans
// }

// 实现2
func reversePrint(head *ListNode) []int {
	count := 0
	for temp := head; temp != nil; temp = temp.Next {
		count++
	}
	ans := make([]int, count)
	for temp := head; temp != nil; temp = temp.Next {
		count--
		ans[count] = temp.Val
	}
	return ans
}
