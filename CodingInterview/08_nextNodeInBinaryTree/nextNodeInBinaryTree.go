package nextNodeInBinaryTree

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode.Right != nil {
		pNode = pNode.Right
		for pNode.Left != nil {
			pNode = pNode.Left
		}
		return pNode
	}
	for pNode.Next != nil {
		if pNode == pNode.Next.Left {
			return pNode.Next
		}
		pNode = pNode.Next
	}
	return nil
}
