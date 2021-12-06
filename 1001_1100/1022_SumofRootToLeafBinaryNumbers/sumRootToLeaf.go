package sumRootToLeaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	var res, sum = 0, 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum = sum<<1 + root.Val
		if root.Left == nil && root.Right == nil {
			res += sum
			return
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}
	dfs(root, sum)
	return res
}
