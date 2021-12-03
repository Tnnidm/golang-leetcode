package isCousins

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	thisLayer := []*TreeNode{root}
	nextLayer := []*TreeNode{}
	var xFound, yFound bool
	var xFather, yFather *TreeNode
	for len(thisLayer) != 0 {
		xFound = false
		yFound = false
		for _, node := range thisLayer {
			if node.Left != nil {
				if node.Left.Val == x {
					xFound = true
					xFather = node
				}
				if node.Left.Val == y {
					yFound = true
					yFather = node
				}
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				if node.Right.Val == x {
					xFound = true
					xFather = node
				}
				if node.Right.Val == y {
					yFound = true
					yFather = node
				}
				nextLayer = append(nextLayer, node.Right)
			}
		}
		if xFound && yFound && xFather != yFather {
			return true
		}
		thisLayer = nextLayer
		nextLayer = []*TreeNode{}
	}
	return false
}
