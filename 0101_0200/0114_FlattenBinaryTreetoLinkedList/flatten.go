package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 迭代方法
// func flatten(root *TreeNode) {
// 	dummyHead := &TreeNode{}
// 	temp := dummyHead
// 	stack := [](*TreeNode){}
// 	stackLen := 0
// 	for root != nil || len(stack) != 0 {
// 		if root == nil {
// 			stackLen--
// 			root = stack[stackLen]
// 			stack = stack[:stackLen]
// 		}
// 		temp.Right = root
// 		temp.Left = nil
// 		temp = temp.Right
// 		if root.Right != nil {
// 			stack = append(stack, root.Right)
// 			stackLen++
// 		}
// 		root = root.Left
// 	}
// 	root = dummyHead.Right
// }

// 递归方法
func flatten(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	ptr := root
	for ptr.Right != nil {
		ptr = ptr.Right
	}
	ptr.Right = temp
}
