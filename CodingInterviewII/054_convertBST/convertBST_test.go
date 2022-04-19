package convertBST

import (
	"reflect"
	"testing"
)

func Test_convertBST(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect *TreeNode
	}{
		{"case 1", &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{3, nil, nil}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := convertBST(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
