package maxPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := root.Val
	var maxSubPathSum func(root *TreeNode) int
	maxSubPathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSubTreeSum := maxSubPathSum(root.Left)
		rightSubTreeSum := maxSubPathSum(root.Right)
		result = max(result, root.Val+max(0, leftSubTreeSum)+max(0, rightSubTreeSum))
		return root.Val + max(0, max(leftSubTreeSum, rightSubTreeSum))
	}
	maxSubPathSum(root)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
