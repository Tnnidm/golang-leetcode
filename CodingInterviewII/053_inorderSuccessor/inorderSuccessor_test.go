package inorderSuccessor

import (
	"reflect"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	p := &TreeNode{1, nil, nil}
	cases := []struct {
		name   string
		input1 *TreeNode
		input2 *TreeNode
		expect *TreeNode
	}{
		{"case 1", p, p, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := inorderSuccessor(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
