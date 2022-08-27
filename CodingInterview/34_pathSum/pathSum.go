package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	result := [][]int{}
	path := []int{}
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		target -= root.Val
		if root.Left == nil && root.Right == nil {
			if target == 0 {
				result = append(result, append([]int{}, path...))
			}
		} else {
			dfs(root.Left, target)
			dfs(root.Right, target)
		}
		path = path[:len(path)-1]
	}
	dfs(root, target)
	return result
}
