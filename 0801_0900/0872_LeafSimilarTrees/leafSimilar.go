package leafSimilar

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := getLeaf(root1)
	leaf2 := getLeaf(root2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}

func getLeaf(root *TreeNode) []int {
	result := []int{}
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root.Left == nil && root.Right == nil {
		*result = append(*result, root.Val)
	} else {
		if root.Left != nil {
			dfs(root.Left, result)
		}
		if root.Right != nil {
			dfs(root.Right, result)
		}
	}
}
