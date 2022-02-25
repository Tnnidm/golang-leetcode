package myPow

import (
	"math"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input1 float64
		input2 int
		expect float64
	}{
		// input1有可能是正数，1，0，负数
		// input2也可能是正数，1，0，负数
		{"case 1.1", 2.00000, 10, 1024.00000},
		{"case 1.2", 2.00000, 0, 1},
		{"case 1.3", 2.00000, 1, 2},
		{"case 1.4", 2.00000, -2, 0.25000},

		{"case 2.1", 1, 10, 1},
		{"case 2.2", 1, 1, 1},
		{"case 2.3", 1, 0, 1},
		{"case 2.4", 1, -2, 1},

		{"case 3.1", 0, 10, 0},
		{"case 3.2", 0, 1, 0},
		{"case 3.3", 0, 0, 1},
		{"case 3.4", 0, -2, math.Inf(0)},

		{"case 4.1", -2.00000, 10, 1024.00000},
		{"case 4.2", -2.00000, 9, -512.0000},
		{"case 4.3", -2.00000, 1, -2},
		{"case 4.4", -2.00000, -2, 0.25000},
		{"case 4.5", -2.00000, -1, -0.5000},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := myPow(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
