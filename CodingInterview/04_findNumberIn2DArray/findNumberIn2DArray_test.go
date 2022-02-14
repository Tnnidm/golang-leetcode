package findNumberIn2DArray

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		matrix [][]int
		target int
		expect bool
	}{
		{"test case 1", [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
		{"test case 2", [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
		{"test case 3", [][]int{{}}, 1, false},
		{"test case 4", [][]int{{-1, 3}}, -1, true},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := findNumberIn2DArray(c.matrix, c.target)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.matrix, c.target)
				}
			},
		)
	}
}
