package largestValues

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := [](*TreeNode){root}
	queueLen := 1
	for queueLen != 0 {
		maxVal := queue[0].Val
		for i := 0; i < queueLen; i++ {
			temp := queue[0] // 这里在循环内边出边进可以大幅减少内存消耗
			queue = queue[1:]
			if temp.Val > maxVal {
				maxVal = temp.Val
			}
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		result = append(result, maxVal)
		queueLen = len(queue)
	}
	return result
}
