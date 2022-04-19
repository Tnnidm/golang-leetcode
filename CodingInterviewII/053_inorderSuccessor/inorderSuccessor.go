package inorderSuccessor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var firstParent *TreeNode
	for p.Val != root.Val {
		if p.Val < root.Val {
			firstParent = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	if root.Right == nil {
		return firstParent
	}
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root
}
