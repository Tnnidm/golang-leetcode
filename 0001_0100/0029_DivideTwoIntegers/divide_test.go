package divide

import (
	"math"
	"reflect"
	"testing"
)

func Test_divide(t *testing.T) {
	cases := []struct {
		name     string
		dividend int
		divisor  int
		expect   int
	}{
		{"case 1", 10, 3, 3},
		{"case 2", -10, 3, -3},
		{"case 3", -10, -3, 3},
		{"case 4", -10, 3, -3},
		{"case 5", 0, 3, 0},
		{"case 6", 100000000000, 1, math.MaxInt32},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := divide(c.dividend, c.divisor)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.dividend, c.divisor)
			}
		})
	}
}
