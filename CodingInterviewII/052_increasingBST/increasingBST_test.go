package increasingBST

import (
	"reflect"
	"testing"
)

func Test_increasingBST(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect *TreeNode
	}{
		{"case 1", &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := increasingBST(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
