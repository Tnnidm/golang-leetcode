package convertBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	lastValue := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Right)
			root.Val += lastValue
			lastValue = root.Val
			dfs(root.Left)
		}
	}
	dfs(root)
	return root
}
