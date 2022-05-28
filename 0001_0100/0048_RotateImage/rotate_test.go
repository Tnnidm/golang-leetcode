package rotate

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{"case 1", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rotate(c.input)
			if !reflect.DeepEqual(c.input, c.expect) {
				t.Fatalf("expect: %v, but got: %v", c.expect, c.input)
			}
		})
	}
}
