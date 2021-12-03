package isUnivalTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		if isUnivalTree(root.Left) && isUnivalTree(root.Right) {
			result := true
			if root.Left != nil {
				result = result && root.Val == root.Left.Val
			}
			if root.Right != nil {
				result = result && root.Val == root.Right.Val
			}
			return result
		} else {
			return false
		}
	}
}
