package serializeBinaryTree

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serializeBinaryTree(root *TreeNode) []string {
	result := []string{}
	if root == nil {
		return result
	}
	serializeDFS(root, &result)
	return result
}

func serializeDFS(root *TreeNode, result *[]string) {
	if root == nil {
		*result = append(*result, "null")
	} else {
		*result = append(*result, strconv.Itoa(root.Val))
		serializeDFS(root.Left, result)
		serializeDFS(root.Right, result)
	}
}

func deserializeBinaryTree(serial []string) *TreeNode {
	if len(serial) == 0 {
		return nil
	}
	root := &TreeNode{0, nil, nil}
	index := -1
	deserializeDFS(root, true, &serial, &index)
	return root.Left
}

func deserializeDFS(root *TreeNode, left bool, serial *[]string, index *int) {
	(*index)++
	if (*serial)[*index] != "null" {
		rootValue, _ := strconv.Atoi((*serial)[*index])
		if left {
			root.Left = &TreeNode{rootValue, nil, nil}
			deserializeDFS(root.Left, true, serial, index)
			deserializeDFS(root.Left, false, serial, index)
		} else {
			root.Right = &TreeNode{rootValue, nil, nil}
			deserializeDFS(root.Right, true, serial, index)
			deserializeDFS(root.Right, false, serial, index)
		}
	}
}
