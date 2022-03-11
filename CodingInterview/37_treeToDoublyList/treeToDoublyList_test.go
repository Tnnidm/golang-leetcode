package treeToDoublyList

import (
	"reflect"
	"testing"
)

func getDoublyList(head *TreeNode) []int {
	result := []int{}
	for p := head; p != nil; p = p.Right {
		result = append(result, p.Val)
	}
	return result
}
func TestTreeToDoublyList(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect []int
	}{
		{"case 1", &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}, []int{1, 2, 3, 4, 5}},
		{"case 2", &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{5, nil, nil}}, []int{1, 2, 4, 5}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := treeToDoublyList(c.input)
			result := getDoublyList(ret)
			if !reflect.DeepEqual(result, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, result, c.input)
			}
		})
	}
}
