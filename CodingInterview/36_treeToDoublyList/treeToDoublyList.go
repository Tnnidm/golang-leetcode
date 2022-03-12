package treeToDoublyList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	start, _ := FromTreeToDoublyList(root)
	return start
}

// func FromTreeToDoublyList(root *TreeNode) (*TreeNode, *TreeNode) {
// 	if root == nil {
// 		return nil, nil
// 	}
// 	left, right := root, root
// 	if root.Left != nil {
// 		leftStart, leftEnd := FromTreeToDoublyList(root.Left)
// 		leftEnd.Right = root
// 		root.Left = leftEnd
// 		left = leftStart
// 	}
// 	if root.Right != nil {
// 		rightStart, rightEnd := FromTreeToDoublyList(root.Right)
// 		root.Right = rightStart
// 		rightStart.Left = root
// 		right = rightEnd
// 	}
// 	return left, right
// }

func FromTreeToDoublyList(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	leftStart, leftEnd := FromTreeToDoublyList(root.Left)
	rightStart, rightEnd := FromTreeToDoublyList(root.Right)
	if leftEnd != nil {
		leftEnd.Right = root
		root.Left = leftEnd
	} else {
		leftStart = root
	}
	if rightStart != nil {
		root.Right = rightStart
		rightStart.Left = root
	} else {
		rightEnd = root
	}
	return leftStart, rightEnd
}
