package add

import (
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	cases := []struct {
		name   string
		input1 int
		input2 int
		expect int
	}{
		{"case 1", 1, 1, 2},
		{"case 2", 10, 17, 27},
		{"case 3", 0, 100, 100},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := add(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
