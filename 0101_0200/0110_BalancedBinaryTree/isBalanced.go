package isBalanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	result := isBalancedTree(root)
	return (result != -1)
}

func isBalancedTree(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		leftHeight := isBalancedTree(root.Left)
		rightHeight := isBalancedTree(root.Right)
		if leftHeight != -1 && rightHeight != -1 {
			maxHeight, minHeight := maxmin(leftHeight, rightHeight)
			difference := maxHeight - minHeight
			if difference <= 1 {
				return maxHeight + 1
			} else {
				return -1
			}
		} else {
			return -1
		}
	}
}

func maxmin(num1 int, num2 int) (int, int) {
	if num1 >= num2 {
		return num1, num2
	} else {
		return num2, num1
	}
}
