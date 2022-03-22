package maxArea

import (
	"reflect"
	"testing"
)

func Test_maxArea(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"case 2", []int{1, 1}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxArea(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
