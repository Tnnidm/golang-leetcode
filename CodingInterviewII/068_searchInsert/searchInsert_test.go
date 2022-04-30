package searchInsert

import (
	"reflect"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect int
	}{
		{"case 1", []int{1, 3, 5, 6}, 5, 2},
		{"case 2", []int{1, 3, 5, 6}, 2, 1},
		{"case 3", []int{1, 3, 5, 6}, 7, 4},
		{"case 4", []int{1, 3, 5, 6}, 0, 0},
		{"case 5", []int{}, 0, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := searchInsert(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
