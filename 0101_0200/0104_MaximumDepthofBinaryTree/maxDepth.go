package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
}
