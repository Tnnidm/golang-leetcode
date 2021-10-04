package minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // My method
// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	} else {
// 		leftDepth := minDepth(root.Left)
// 		rightDepth := minDepth(root.Right)
// 		if leftDepth != 0 && rightDepth != 0 {
// 			return min(leftDepth, rightDepth) + 1
// 		} else {
// 			return max(leftDepth, rightDepth) + 1
// 		}
// 	}
// }

// func min(num1 int, num2 int) int {
// 	if num1 <= num2 {
// 		return num1
// 	} else {
// 		return num2
// 	}
// }

// func max(num1 int, num2 int) int {
// 	if num1 >= num2 {
// 		return num1
// 	} else {
// 		return num2
// 	}
// }

// method 2: BFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// A BFS algorithm
	row := []*TreeNode{root}
	depth := 0
	for len(row) > 0 {
		depth++
		nextRow := []*TreeNode{}
		for i := 0; i < len(row); i++ {
			node := row[i]
			if node.Left != nil {
				nextRow = append(nextRow, node.Left)
			}
			if node.Right != nil {
				nextRow = append(nextRow, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return depth
			}
		}
		row = nextRow
	}
	return depth
}
