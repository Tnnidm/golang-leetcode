package trap

import (
	"reflect"
	"testing"
)

func Test_trap(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"case 2", []int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := trap(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
