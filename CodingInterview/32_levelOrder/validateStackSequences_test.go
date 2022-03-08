package levelOrder

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	cases := []struct {
		name    string
		input   *TreeNode
		expect1 []int
		expect2 [][]int
		expect3 [][]int
	}{
		{"case 1", nil, []int{}, [][]int{}, [][]int{}},
		{"case 1", &TreeNode{0, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, []int{0, 1, 2}, [][]int{{0}, {1, 2}}, [][]int{{0}, {2, 1}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret1 := levelOrderPlain(c.input)
			if !reflect.DeepEqual(ret1, c.expect1) {
				t.Fatalf("in levelOrderPlain, expect: %v, but got: %v, with input: %v", c.expect1, ret1, c.input)
			}
			ret2 := levelOrderHierarchy(c.input)
			if !reflect.DeepEqual(ret2, c.expect2) {
				t.Fatalf("in levelOrderHierarchy, expect: %v, but got: %v, with input: %v", c.expect2, ret2, c.input)
			}
			ret3 := levelOrderZType(c.input)
			if !reflect.DeepEqual(ret3, c.expect3) {
				t.Fatalf("in levelOrderZType, expect: %v, but got: %v, with input: %v", c.expect3, ret3, c.input)
			}
		})
	}
}
