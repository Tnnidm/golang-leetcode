package cuttingRope

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{"case 1", 2, 1},
		{"case 2", 3, 2},
		{"case 3", 4, 4},
		{"case 4", 5, 6},
		{"case 5", 10, 36},
		{"case 6", 58, 1549681956},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := cuttingRope(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
