package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	pathMap := map[int]int{}
	num := 0
	preOrder(root, 0, targetSum, &pathMap, &num)
	return num
}

func preOrder(root *TreeNode, path int, targetSum int, pathMap *map[int]int, num *int) {
	if root != nil {
		path += root.Val
		if path == targetSum {
			*num++
		}
		if (*pathMap)[path-targetSum] != 0 {
			*num += (*pathMap)[path-targetSum]
		}
		(*pathMap)[path]++
		preOrder(root.Left, path, targetSum, pathMap, num)
		preOrder(root.Right, path, targetSum, pathMap, num)
		(*pathMap)[path]--
	}
}
