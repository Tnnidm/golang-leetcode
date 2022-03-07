package isSubStructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isASubStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isASubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	} else {
		return isASubStructure(A.Left, B.Left) && isASubStructure(A.Right, B.Right)
	}
}
