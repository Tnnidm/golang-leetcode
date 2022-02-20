package minArray

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 0", []int{1, 2, 3, 4, 5}, 1},
		{"case 1", []int{3, 4, 5, 1, 2}, 1},
		{"case 2", []int{2, 2, 2, 0, 1}, 0},
		{"case 3", []int{2, 2, 2, 2, 2}, 2},
		{"case 4", []int{3, 4, 5, 1, 1}, 1},
		{"case 5", []int{0}, 0},
		{"case 6", []int{3, 3, 1, 3}, 1},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := minArray(c.input)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
				}
			},
		)
	}
}
