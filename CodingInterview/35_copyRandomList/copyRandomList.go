package copyRandomList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// // 时间复杂度O(n)，空间复杂度O(n)
// func copyRandomList(head *Node) *Node {
// 	if head == nil {
// 		return nil
// 	}
// 	count := 1
// 	for p1 := head; p1 != nil; p1 = p1.Next {
// 		count++
// 	}
// 	newHead := &Node{head.Val, nil, nil}
// 	addressMap := make(map[*Node](*Node), count)
// 	addressMap[head] = newHead
// 	for p1, p2 := head.Next, newHead; p1 != nil; p1, p2 = p1.Next, p2.Next {
// 		p2.Next = &Node{p1.Val, nil, nil}
// 		addressMap[p1] = p2.Next
// 	}
// 	for p1, p2 := head, newHead; p1 != nil; p1, p2 = p1.Next, p2.Next {
// 		p2.Random = addressMap[p1.Random]
// 	}
// 	return newHead
// }

// 时间复杂度O(n)，空间复杂度O(1)
// 但是题目不一定允许修改原始链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for p1 := head; p1 != nil; {
		temp := p1.Next
		p1.Next = &Node{p1.Val, temp, nil}
		p1 = temp
	}
	for p1 := head; p1 != nil; p1 = p1.Next.Next {
		if p1.Random != nil {
			p1.Next.Random = p1.Random.Next
		}
	}
	for p1 := head.Next; p1 != nil; p1 = p1.Next {
		temp := p1.Next
		temp.Next = nil
		p1.Next = p1.Next.Next
	}
	return head.Next
}
