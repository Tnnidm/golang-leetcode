package mirrorTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	mirrorLeftandRight(root)
	return root
}

func mirrorLeftandRight(root *TreeNode) {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		mirrorLeftandRight(root.Left)
		mirrorLeftandRight(root.Right)
	}
}
