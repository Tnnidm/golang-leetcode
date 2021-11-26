package minDiffInBST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	preValue := -1
	minDiff := math.MaxInt32
	diff := 0
	dfs(root, &preValue, &diff, &minDiff)
	return minDiff
}

func dfs(root *TreeNode, preValue *int, diff *int, minDiff *int) {
	if root != nil {
		dfs(root.Left, preValue, diff, minDiff)
		if *preValue != -1 {
			*diff = root.Val - *preValue
			if *diff < *minDiff {
				*minDiff = *diff
			}
		}
		*preValue = root.Val
		dfs(root.Right, preValue, diff, minDiff)
	}
}
