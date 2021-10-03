package sortedArrayToBST

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	// 测试用例
	cases := []struct {
		name   string
		input  []int
		expect *TreeNode
	}{
		{"case 1", []int{1}, &TreeNode{Val: 1}},
		{
			"case 2",
			[]int{-10, -3, 0, 5, 9},
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := sortedArrayToBST(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
