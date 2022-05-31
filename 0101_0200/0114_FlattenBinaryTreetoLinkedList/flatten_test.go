package flatten

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	cases := []struct {
		name   string
		input  *TreeNode
		expect int
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			flatten(c.input)
			if !reflect.DeepEqual(c.input, c.expect) {
				t.Fatalf("expect: %v, but got: %v", c.expect, c.input)
			}
		})
	}
}
