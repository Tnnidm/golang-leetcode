package depthOfBinaryTree

func isBalanced(root *TreeNode) bool {
	ans, _ := dfs(root)
	return ans
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftFlag, leftDepth := dfs(root.Left)
	if !leftFlag {
		return false, 0
	}
	rightFlag, rightDepth := dfs(root.Right)
	if !rightFlag {
		return false, 0
	}
	if leftDepth < rightDepth {
		leftDepth, rightDepth = rightDepth, leftDepth
	}
	if leftDepth <= rightDepth+1 {
		return true, leftDepth + 1
	}
	return false, 0
}
