package buildTree

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 []int
		expect *TreeNode
	}{
		{"case 1", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}},
		{"case 2", []int{3}, []int{3}, &TreeNode{3, nil, nil}},
		{"case 3", []int{}, []int{}, nil},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := buildTree(c.input1, c.input2)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
				}
			},
		)
	}
}
