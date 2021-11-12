package tree2str

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := strconv.Itoa(root.Val)
	if root.Left != nil {
		result += "(" + tree2str(root.Left) + ")"
	} else if root.Right != nil {
		result += "()"
	}
	if root.Right != nil {
		result += "(" + tree2str(root.Right) + ")"
	}
	return result
}
