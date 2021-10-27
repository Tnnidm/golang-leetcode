package binaryTreePaths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	treePath(root, strconv.Itoa(root.Val), &result)
	return result
}

func treePath(root *TreeNode, path string, result *[]string) {
	if root.Left == nil && root.Right == nil {
		*result = append(*result, path)
	} else {
		if root.Left != nil {
			treePath(root.Left, path+"->"+strconv.Itoa(root.Left.Val), result)
		}
		if root.Right != nil {
			treePath(root.Right, path+"->"+strconv.Itoa(root.Right.Val), result)
		}
	}
}
