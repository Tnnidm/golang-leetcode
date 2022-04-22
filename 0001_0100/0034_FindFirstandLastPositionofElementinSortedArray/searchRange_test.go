package searchRange

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect []int
	}{
		{"case 1", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"case 2", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"case 3", []int{}, 8, []int{-1, -1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := searchRange(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
