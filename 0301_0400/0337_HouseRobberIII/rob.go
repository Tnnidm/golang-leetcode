package rob

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	includeRoot, excludeRoot := robCore(root)
	return max(includeRoot, excludeRoot)
}

func robCore(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val, 0
	}
	includeLeft, excludeLeft := robCore(root.Left)
	includeRight, excludeRight := robCore(root.Right)
	return root.Val + excludeLeft + excludeRight, max(includeLeft, excludeLeft) + max(includeRight, excludeRight)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
