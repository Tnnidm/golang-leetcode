package findBottomLeftValue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 使用队列解决
// func findBottomLeftValue(root *TreeNode) int {
// 	queue := [](*TreeNode){root}
// 	var layerFirst int
// 	queueLen := 1
// 	for queueLen != 0 {
// 		layerFirst = queue[0].Val
// 		for i := 0; i < queueLen; i++ {
// 			if queue[0].Left != nil {
// 				queue = append(queue, queue[0].Left)
// 			}
// 			if queue[0].Right != nil {
// 				queue = append(queue, queue[0].Right)
// 			}
// 			queue = queue[1:]
// 		}
// 		queueLen = len(queue)
// 	}
// 	return layerFirst
// }

// 使用栈解决
// 利用中序遍历优先遍历一层的最左边的特点
func findBottomLeftValue(root *TreeNode) int {
	var val, maxDepth int
	dfs(root, 1, &val, &maxDepth)
	return val
}

func dfs(root *TreeNode, depth int, val *int, maxDepth *int) {
	if root != nil {
		dfs(root.Left, depth+1, val, maxDepth)
		if depth > *maxDepth {
			*val = root.Val
			*maxDepth = depth
		}
		dfs(root.Right, depth+1, val, maxDepth)
	}
}
