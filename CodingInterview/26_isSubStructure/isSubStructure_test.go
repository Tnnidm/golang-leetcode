package isSubStructure

import (
	"reflect"
	"testing"
)

func Test_isSubStructure(t *testing.T) {
	cases := []struct {
		name   string
		input1 *TreeNode
		input2 *TreeNode
		expect bool
	}{
		{"case 1", nil, nil, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isSubStructure(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
