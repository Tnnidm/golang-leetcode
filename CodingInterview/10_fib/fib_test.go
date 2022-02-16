package fib

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
		{"case 1", 0, 0},
		{"case 2", 1, 1},
		{"case 3", 2, 1},
		{"case 4", 3, 2},
		{"case 5", 100, 687995182},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := fib(c.input)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
				}
			},
		)
	}
}
