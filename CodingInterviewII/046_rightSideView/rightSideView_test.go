package rightSideView

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect []int
	}{
		{"case 1", &TreeNode{1, &TreeNode{2, &TreeNode{5, nil, nil}, nil}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}, []int{1, 3, 4}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := rightSideView(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
