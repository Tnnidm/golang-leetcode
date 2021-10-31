package sumOfLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left != nil && isALeave(root.Left) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	} else {
		return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	}
}

func isALeave(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	} else {
		return false
	}
}
