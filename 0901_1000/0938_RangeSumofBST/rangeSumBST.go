package rangeSumBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	dfs(root, &low, &high, &sum)
	return sum
}

func dfs(root *TreeNode, low *int, high *int, sum *int) {
	if root != nil {
		if root.Val > *low {
			dfs(root.Left, low, high, sum)
		}
		if root.Val <= *high && root.Val >= *low {
			*sum += root.Val
		}
		if root.Val < *high {
			dfs(root.Right, low, high, sum)
		}
	}
}
