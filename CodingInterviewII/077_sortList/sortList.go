package sortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	half := split(head)
	head = sortList(head)
	half = sortList(half)
	return mergeList(head, half)
}

func split(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	half := slow.Next
	slow.Next = nil
	return half
}

func mergeList(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	} else {
		cur.Next = head2
	}
	return dummyHead.Next
}
