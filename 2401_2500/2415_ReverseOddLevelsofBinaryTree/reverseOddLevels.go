package reverseOddLevels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	oddFlag := false
	queue := [](*TreeNode){root}
	queueLen := 1
	for queueLen != 0 {
		if oddFlag {
			for i := 0; i < queueLen>>1; i++ {
				queue[i].Val, queue[queueLen-1-i].Val = queue[queueLen-1-i].Val, queue[i].Val
			}
		}
		for i := 0; i < queueLen; i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		oddFlag = !oddFlag
		queueLen = len(queue)
	}
	return root
}
