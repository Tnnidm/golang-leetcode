package insert

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	a := &Node{1, nil}
	a.Next = a
	cases := []struct {
		name   string
		input1 *Node
		input2 int
		expect *Node
	}{
		{"case 1", nil, 1, a},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := insert(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
