package search

import (
	"reflect"
	"testing"
)

func Test_search(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect int
	}{
		{"case 1", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"case 2", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"case 1", []int{1}, 0, -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := search(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
