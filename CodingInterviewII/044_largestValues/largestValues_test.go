package largestValues

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect []int
	}{
		{"case 1", &TreeNode{1, &TreeNode{3, &TreeNode{5, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{9, nil, nil}}}, []int{1, 2, 9}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := largestValues(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
