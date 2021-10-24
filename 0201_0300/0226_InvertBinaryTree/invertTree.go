package invertTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	} else {
		root.Left, root.Right = root.Right, root.Left
		root.Left = invertTree(root.Left)
		root.Right = invertTree(root.Right)
		return root
	}
}
