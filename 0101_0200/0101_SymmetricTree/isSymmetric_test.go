package isSymmetric

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	// 测试用例
	cases := []struct {
		name   string
		input  *TreeNode
		expect bool
	}{
		{"case 1", &TreeNode{Val: 1}, true},
		{
			"case 2",
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			true,
		},
		{
			"case 3",
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isSymmetric(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
