package levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := [](*TreeNode){root}
	queueLen := 1
	result := [][]int{}
	for queueLen != 0 {
		layerValue := []int{}
		for i := 0; i < queueLen; i++ {
			node := queue[0]
			queue = queue[1:]
			layerValue = append(layerValue, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, layerValue)
		queueLen = len(queue)
	}
	return result
}
