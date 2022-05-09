package minCostClimbingStairs

import (
	"reflect"
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{10, 15, 20}, 15},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minCostClimbingStairs(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
