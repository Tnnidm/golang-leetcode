package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	pathMap := map[int]int{0: 1}
	result := 0
	var preOrder func(root *TreeNode, sum int)
	preOrder = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		result += pathMap[sum-targetSum]
		pathMap[sum]++
		preOrder(root.Left, sum)
		preOrder(root.Right, sum)
		pathMap[sum]--
	}
	preOrder(root, 0)
	return result
}
