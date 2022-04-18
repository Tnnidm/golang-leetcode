package maxPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	sum := root.Val
	maxSubPath(root, &sum)
	return sum
}

func maxSubPath(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	leftPathSum := maxSubPath(root.Left, maxSum)
	rightPathSum := maxSubPath(root.Right, maxSum)
	pathSum := root.Val
	if leftPathSum > 0 {
		pathSum += leftPathSum
	}
	if rightPathSum > 0 {
		pathSum += rightPathSum
	}
	*maxSum = max(*maxSum, pathSum)
	if max(leftPathSum, rightPathSum) < 0 {
		return root.Val
	}
	return root.Val + max(leftPathSum, rightPathSum)
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
