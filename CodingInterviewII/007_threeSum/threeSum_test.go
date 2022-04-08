package threeSum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{"case 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"case 2", []int{}, [][]int{}},
		{"case 3", []int{0}, [][]int{}},
		{"case 4", []int{0, 0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := threeSum(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
