package maxPathSum

import (
	"reflect"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect int
	}{
		{"case 1", &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 6},
		{"case 2", &TreeNode{-10, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}, 42},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxPathSum(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
