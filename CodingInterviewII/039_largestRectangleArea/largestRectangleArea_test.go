package largestRectangleArea

import (
	"reflect"
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"case 2", []int{5, 5, 5}, 15},
		{"case 3", []int{2, 4}, 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := largestRectangleArea(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
