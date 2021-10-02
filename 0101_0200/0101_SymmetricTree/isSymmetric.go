package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
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
			} else if !isSymmetricTree(p.Left, q.Right) {
				return false
			} else if !isSymmetricTree(p.Right, q.Left) {
				return false
			}
			return true
		}
	}
}
