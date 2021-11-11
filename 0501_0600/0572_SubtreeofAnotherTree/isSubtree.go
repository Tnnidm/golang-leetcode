package isSubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		} else {
			return false
		}
	} else {
		if q == nil {
			return false
		} else {
			if p.Val != q.Val {
				return false
			}
			if !isSameTree(p.Left, q.Left) {
				return false
			}
			if !isSameTree(p.Right, q.Right) {
				return false
			}
			return true
		}
	}
}
