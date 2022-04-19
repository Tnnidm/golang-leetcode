package increasingBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	dummyHead := &TreeNode{}
	temp := &dummyHead
	dfs(root, temp)
	return dummyHead.Right
}

func dfs(root *TreeNode, dummyHead **TreeNode) {
	if root != nil {
		dfs(root.Left, dummyHead)
		(*dummyHead).Right = root
		root.Left = nil
		(*dummyHead) = root
		dfs(root.Right, dummyHead)
	}
}
