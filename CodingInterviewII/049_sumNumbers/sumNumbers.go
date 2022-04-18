package sumNumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	dfs(root, 0, &sum)
	return sum
}

func dfs(root *TreeNode, layerSum int, sum *int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*sum += 10*layerSum + root.Val
			return
		}
		dfs(root.Left, 10*layerSum+root.Val, sum)
		dfs(root.Right, 10*layerSum+root.Val, sum)
	}
}
