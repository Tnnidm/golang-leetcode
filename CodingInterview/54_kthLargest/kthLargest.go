package kthLargest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var ans int
	dfs(root, &k, &ans)
	return ans
}

func dfs(root *TreeNode, k *int, ans *int) {
	if root.Right != nil {
		dfs(root.Right, k, ans)
	}
	*k--
	if *k == 0 {
		*ans = root.Val
	}
	if *k > 0 && root.Left != nil {
		dfs(root.Left, k, ans)
	}
}
