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
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := lowestCommonAncestor(c.root, c.p, c.q)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v, %v, %v", c.expect, ret, c.root, c.p, c.q)
			}
		})
	}
}
