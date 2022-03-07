package spiralOrder

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect []int
	}{
		{"case 1", [][]int{}, []int{}},
		{"case 2", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := spiralOrder(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
