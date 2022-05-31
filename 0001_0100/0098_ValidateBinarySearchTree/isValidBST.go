package isValidBST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isValidBST(root *TreeNode) bool {
// 	lastNum := math.MinInt32
// 	first := true
// 	result := true
// 	var dfs func(root *TreeNode)
// 	dfs = func(root *TreeNode) {
// 		if root != nil && result {
// 			dfs(root.Left)
// 			if root.Val > lastNum || first {
// 				lastNum = root.Val
// 				first = false
// 			} else {
// 				result = false
// 			}
// 			dfs(root.Right)
// 		}
// 	}
// 	dfs(root)
// 	return result
// }

func isValidBST(root *TreeNode) bool {
	return isValidbst(root, -math.MaxInt64, math.MaxInt64)
}

func isValidbst(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && isValidbst(root.Left, min, root.Val) && isValidbst(root.Right, root.Val, max)
}
