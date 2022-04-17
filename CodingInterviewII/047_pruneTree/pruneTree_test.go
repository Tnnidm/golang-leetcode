package pruneTree

import (
	"reflect"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect *TreeNode
	}{
		{"case 1", &TreeNode{1, &TreeNode{0, &TreeNode{0, nil, nil}, nil}, &TreeNode{1, nil, &TreeNode{0, nil, nil}}}, &TreeNode{1, nil, &TreeNode{1, nil, nil}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := pruneTree(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
