package mincostTickets

import (
	"reflect"
	"testing"
)

func Test_mincostTickets(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 []int
		expect int
	}{
		{"case 1", []int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
		{"case 2", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
		{"case 3", []int{6, 8, 9, 18, 20, 21, 23, 25}, []int{2, 10, 41}, 16},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := mincostTickets(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
