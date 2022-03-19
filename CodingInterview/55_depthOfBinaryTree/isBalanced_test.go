package depthOfBinaryTree

import (
	"reflect"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect bool
	}{
		{"case 1", nil, true},
		{"case 2", &TreeNode{1, nil, nil}, true},
		{"case 3", &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isBalanced(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
