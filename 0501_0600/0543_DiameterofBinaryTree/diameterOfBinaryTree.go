package diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max_length := 0
	diameter(root, &max_length)
	return max_length
}

func diameter(root *TreeNode, max_length *int) int {
	if root == nil {
		return 0
	} else {
		left_depth := diameter(root.Left, max_length)
		right_depth := diameter(root.Right, max_length)
		if left_depth+right_depth > *max_length {
			*max_length = left_depth + right_depth
		}
		if left_depth > right_depth {
			return left_depth + 1
		} else {
			return right_depth + 1
		}
	}
}
