package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	result := [][]int{}
	dfs(root, target, &result, []int{})
	return result
}

func dfs(root *TreeNode, target int, result *[][]int, path []int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	target -= root.Val
	if root.Left == nil && root.Right == nil {
		if target == 0 {
			pathLen := len(path)
			newpath := make([]int, pathLen)
			for i := 0; i < pathLen; i++ {
				newpath[i] = path[i]
			}
			(*result) = append((*result), newpath)
		}
	} else {
		dfs(root.Left, target, result, path)
		dfs(root.Right, target, result, path)
	}
}
