package postorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorder(root, &result)
	return result
}

func postorder(root *TreeNode, result *[]int) {
	if root != nil {
		postorder(root.Left, result)
		postorder(root.Right, result)
		*result = append(*result, root.Val)
	}
}
