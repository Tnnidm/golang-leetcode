package serializeBinaryTree

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
func (codec *Codec) serialize(root *TreeNode) string {
	result := strings.Builder{}
	queue := [](*TreeNode){root}
	queueLen := 1
	for queueLen != 0 {
		temp := queue[0]
		queue = queue[1:]
		queueLen--
		if temp == nil {
			result.WriteByte(',')
			result.WriteByte('#')
		} else {
			result.WriteByte(',')
			result.WriteString(strconv.Itoa(temp.Val))
			queue = append(queue, temp.Left)
			queue = append(queue, temp.Right)
			queueLen += 2
		}
	}
	return result.String()[1:]
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	dataList := strings.Split(data, ",")
	root := transfer(dataList[0])
	if root == nil {
		return root
	}
	queue := [](*TreeNode){root}
	for i := 1; i < len(dataList); {
		temp := queue[0]
		queue = queue[1:]
		temp.Left = transfer(dataList[i])
		i++
		temp.Right = transfer(dataList[i])
		i++
		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}
	}
	return root
}

func transfer(node string) *TreeNode {
	if node == "#" {
		return nil
	}
	num, _ := strconv.Atoi(node)
	return &TreeNode{num, nil, nil}
}
