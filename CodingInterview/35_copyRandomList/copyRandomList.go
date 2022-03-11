package copyRandomList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// // assistNode应该可以不要
// func copyRandomList(head *Node) *Node {
// 	assistHead := &Node{}
// 	assistHeadTemp := assistHead
// 	listMap := map[*Node](*Node){nil: nil}
// 	for temp := head; temp != nil; temp = temp.Next {
// 		assistHeadTemp.Next = &Node{temp.Val, nil, nil}
// 		listMap[temp] = assistHeadTemp.Next
// 		assistHeadTemp = assistHeadTemp.Next
// 	}
// 	for p1, p2 := head, assistHead.Next; p1 != nil; p1, p2 = p1.Next, p2.Next {
// 		p2.Random = listMap[p1.Random]
// 	}
// 	return assistHead.Next
// }

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	count := 1
	for p1 := head; p1 != nil; p1 = p1.Next {
		count++
	}
	newHead := &Node{head.Val, nil, nil}
	addressMap := make(map[*Node](*Node), count)
	addressMap[head] = newHead
	for p1, p2 := head.Next, newHead; p1 != nil; p1, p2 = p1.Next, p2.Next {
		p2.Next = &Node{p1.Val, nil, nil}
		addressMap[p1] = p2.Next
	}
	for p1, p2 := head, newHead; p1 != nil; p1, p2 = p1.Next, p2.Next {
		p2.Random = addressMap[p1.Random]
	}
	return newHead
}
