package maxPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := root.Val
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftPathSum := max(0, dfs(root.Left))
		rightPathSum := max(0, dfs(root.Right))
		result = max(result, leftPathSum+root.Val+rightPathSum)
		return root.Val + max(leftPathSum, rightPathSum)
	}
	dfs(root)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
