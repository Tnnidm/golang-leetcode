package rightSideView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := [](*TreeNode){root}
	queueLen := 1
	for queueLen != 0 {
		result = append(result, queue[0].Val)
		for i := 0; i < queueLen; i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
		}
		queueLen = len(queue)
	}
	return result
}
