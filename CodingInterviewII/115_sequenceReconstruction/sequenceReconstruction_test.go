package sequenceReconstruction

import (
	"reflect"
	"testing"
)

func Test_sequenceReconstruction(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 [][]int
		expect bool
	}{
		{"case 1", []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}, false},
		{"case 2", []int{1, 2, 3}, [][]int{{1, 2}}, false},
		{"case 3", []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}, true},
		{"case 4", []int{4, 1, 5, 2, 6, 3}, [][]int{{5, 2, 6, 3}, {4, 1, 5, 2}}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := sequenceReconstruction(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
