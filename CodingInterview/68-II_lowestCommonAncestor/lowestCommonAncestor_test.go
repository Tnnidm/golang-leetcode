package lowestCommonAncestor

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	cases := []struct {
		name   string
		root   *TreeNode
		p      *TreeNode
		q      *TreeNode
		expect *TreeNode
	}{
		{"case 1", nil, nil, nil, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := lowestCommonAncestor(c.root, c.p, c.q)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v", c.expect, ret)
			}
		})
	}

}
