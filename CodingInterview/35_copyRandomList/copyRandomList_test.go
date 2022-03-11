package copyRandomList

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	cases := []struct {
		name   string
		input  *Node
		expect *Node
	}{
		{"case 1", nil, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := copyRandomList(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
