package isSameTree

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	// 测试用例
	cases := []struct {
		name   string
		input1 *TreeNode
		input2 *TreeNode
		expect bool
	}{
		{"case 1", &TreeNode{Val: 1}, &TreeNode{Val: 1}, true},
		{"case 2", nil, nil, true},
		{
			"case 3",
			&TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
				},
			},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: nil,
			},
			false,
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
			true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isSameTree(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
