package mySqrt

import (
	"math"
	"reflect"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{"case 1", 4, 2},
		{"case 2", 8, 2},
		{"case 3", math.MaxInt32, int(math.Sqrt(math.MaxInt32))},
		{"case 4", 0, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := mySqrt(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
