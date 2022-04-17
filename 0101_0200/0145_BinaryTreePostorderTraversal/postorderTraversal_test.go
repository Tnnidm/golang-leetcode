package postorderTraversal

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	// 测试用例
	cases := []struct {
		name   string
		input  *TreeNode
		expect []int
	}{
		{"case 1", &TreeNode{Val: 1}, []int{1}},
		{"case 2", nil, []int{}},
		{
			"case 3",
			&TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: nil,
				},
			},
			[]int{3, 2, 1},
		},
		{
			"case 4",
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 7},
				},
			},
			[]int{1, 3, 2, 5, 7, 6, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := postorderTraversal(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
