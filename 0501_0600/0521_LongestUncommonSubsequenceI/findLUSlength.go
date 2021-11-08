package findLUSlength

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	pre_value := math.MinInt32
	min_diff := math.MaxInt32
	midOrderSearch(root, &pre_value, &min_diff)
	return min_diff
}

func midOrderSearch(root *TreeNode, pre_value *int, min_diff *int) {
	if root != nil {
		midOrderSearch(root.Left, pre_value, min_diff)
		*min_diff = min(*min_diff, root.Val-*pre_value)
		*pre_value = root.Val
		midOrderSearch(root.Right, pre_value, min_diff)
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
