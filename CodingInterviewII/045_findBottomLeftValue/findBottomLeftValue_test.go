package findBottomLeftValue

import (
	"reflect"
	"testing"
)

func Test_findBottomLeftValue(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect int
	}{
		{"case 1", &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findBottomLeftValue(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
