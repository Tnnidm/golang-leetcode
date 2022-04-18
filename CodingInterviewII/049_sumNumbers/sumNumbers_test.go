package sumNumbers

import (
	"reflect"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect int
	}{
		{"case 1", &TreeNode{1, &TreeNode{0, &TreeNode{0, nil, nil}, nil}, &TreeNode{1, nil, &TreeNode{0, nil, nil}}}, 210},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := sumNumbers(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
