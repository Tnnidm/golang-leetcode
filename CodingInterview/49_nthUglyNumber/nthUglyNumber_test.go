package nthUglyNumber

import (
	"reflect"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{"case 1", 10, 12},
		{"case 2", 1690, 2123366400},
		{"case 3", 1, 1},
		{"case 4", 2, 2},
		{"case 5", 3, 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := nthUglyNumber(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
