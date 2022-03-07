package mirrorTree

import (
	"reflect"
	"testing"
)

func Test_mirrorTree(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect *TreeNode
	}{
		{"case 1", nil, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := mirrorTree(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
