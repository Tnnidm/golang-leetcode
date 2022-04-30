package peakIndexInMountainArray

import (
	"reflect"
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{0, 1, 0}, 1},
		{"case 2", []int{1, 3, 5, 4, 2}, 2},
		{"case 3", []int{1, 10, 7, 4, 2}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := peakIndexInMountainArray(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
