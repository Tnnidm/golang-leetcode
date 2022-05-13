package minCost

import (
	"reflect"
	"testing"
)

func Test_minCost(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{"case 1", [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}, 10},
		{"case 2", [][]int{{7, 6, 2}}, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minCost(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
