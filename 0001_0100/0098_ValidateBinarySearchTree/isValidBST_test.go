package isValidBST

import (
	"reflect"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect *TreeNode
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isValidBST(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
