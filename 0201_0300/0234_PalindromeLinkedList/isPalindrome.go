package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	} else {
		numStack := []int{}
		for head != nil {
			numStack = append(numStack, head.Val)
			head = head.Next
		}
		numStackLen := len(numStack)
		for i := 0; i < numStackLen/2+1; i++ {
			if numStack[i] != numStack[numStackLen-1-i] {
				return false
			}
		}
		return true
	}

}
