package increasingBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	var dummy = &TreeNode{Val: 0}
	head := dummy
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		head.Right = root
		root.Left = nil
		head = head.Right
		dfs(root.Right)
	}
	dfs(root)
	return dummy.Right
}
