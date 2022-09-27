package serialize

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := []string{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			result = append(result, "null")
			return
		}
		valStr := strconv.Itoa(root.Val)
		result = append(result, valStr)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	var dfs func(nodes *[]string) *TreeNode
	dfs = func(nodes *[]string) *TreeNode {
		node := (*nodes)[0]
		(*nodes) = (*nodes)[1:]
		if node == "null" {
			return nil
		}
		value, _ := strconv.Atoi(node)
		result := &TreeNode{value, nil, nil}
		result.Left = dfs(nodes)
		result.Right = dfs(nodes)
		return result
	}
	return dfs(&nodes)
}
