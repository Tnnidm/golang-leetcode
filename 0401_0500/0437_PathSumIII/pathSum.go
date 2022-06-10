package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	result := 0
	preSumSet := map[int]int{0: 1}
	var dfs func(root *TreeNode, preSum int)
	dfs = func(root *TreeNode, preSum int) {
		if root != nil {
			preSum += root.Val
			result += preSumSet[preSum-targetSum]
			preSumSet[preSum]++
			dfs(root.Left, preSum)
			dfs(root.Right, preSum)
			preSumSet[preSum]--
		}
	}
	dfs(root, 0)
	return result
}
