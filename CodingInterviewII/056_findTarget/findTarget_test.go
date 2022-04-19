package findTarget

import (
	"reflect"
	"testing"
)

func Test_findTarget(t *testing.T) {
	cases := []struct {
		name   string
		input1 *TreeNode
		input2 int
		expect bool
	}{
		{"case 1", &TreeNode{1, &TreeNode{2, nil, nil}, nil}, 3, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findTarget(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
