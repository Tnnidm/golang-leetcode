package kSmallestPairs

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 []int
		input3 int
		expect [][]int
	}{
		{"case 1", []int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := kSmallestPairs(c.input1, c.input2, c.input3)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v, %v and %v", c.expect, ret, c.input1, c.input2, c.input3)
			}
		})
	}
}
