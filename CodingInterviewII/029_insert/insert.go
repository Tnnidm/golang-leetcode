package insert

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		newNode := &Node{x, nil}
		newNode.Next = newNode
		return newNode
	}
	var maxNode *Node
	flag := true
	for p := aNode; p != aNode || flag; p = p.Next {
		flag = false
		if p.Val <= x && x <= p.Next.Val {
			p.Next = &Node{x, p.Next}
			return aNode
		}
		if p.Val > p.Next.Val {
			maxNode = p
		}
	}
	if maxNode == nil {
		maxNode = aNode
	}
	maxNode.Next = &Node{x, maxNode.Next}
	return aNode
}
