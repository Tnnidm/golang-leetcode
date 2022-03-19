package depthOfBinaryTree

import (
	"reflect"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect int
	}{
		{"case 1", nil, 0},
		{"case 2", &TreeNode{1, nil, nil}, 1},
		{"case 3", &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}, 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxDepth(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
