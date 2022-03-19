package kthLargest

import (
	"reflect"
	"testing"
)

func Test_kthLargest(t *testing.T) {
	cases := []struct {
		name   string
		input1 *TreeNode
		input2 int
		expect int
	}{
		{"case 1", &TreeNode{1, nil, nil}, 1, 1},
		{"case 2", &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}, 4, 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := kthLargest(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
