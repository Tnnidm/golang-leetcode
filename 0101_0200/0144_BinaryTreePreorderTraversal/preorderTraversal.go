package preorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 最直接的解法
// func preorderTraversal(root *TreeNode) []int {
// 	result := []int{}
// 	if root != nil {
// 		result = append(result, root.Val)
// 		result = append(result, preorderTraversal(root.Left)...)
// 		result = append(result, preorderTraversal(root.Right)...)
// 	}
// 	return result
// }

// 优化内存的解法
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorder(root, &result)
	return result
}

func preorder(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		preorder(root.Left, result)
		preorder(root.Right, result)
	}
}
