package searchBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	} else if val > root.Val {
		if root.Right != nil {
			return searchBST(root.Right, val)
		} else {
			return nil
		}
	} else {
		if root.Left != nil {
			return searchBST(root.Left, val)
		} else {
			return nil
		}
	}
}
