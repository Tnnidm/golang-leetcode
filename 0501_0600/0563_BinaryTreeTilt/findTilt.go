package findTilt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	tiltSum := 0
	nodeSum(root, &tiltSum)
	return tiltSum
}

func nodeSum(root *TreeNode, tiltSum *int) int {
	if root == nil {
		return 0
	} else {
		leftNodeSum := nodeSum(root.Left, tiltSum)
		rightNodeSum := nodeSum(root.Right, tiltSum)
		*tiltSum += abs(leftNodeSum - rightNodeSum)
		return leftNodeSum + rightNodeSum + root.Val
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
