package divide

import (
	"math"
	"reflect"
	"testing"
)

func Test_divide(t *testing.T) {
	cases := []struct {
		name   string
		input1 int
		input2 int
		expect int
	}{
		{"case 1", 50, 5, 10},
		{"case 2", 707, 7, 101},
		{"case 3", math.MinInt32, -1, math.MaxInt32},
		{"case 4", math.MinInt32, math.MinInt32, 1},
		{"case 5", math.MinInt32, 1, math.MinInt32},
		{"case 6", 4, math.MinInt32, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := divide(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
