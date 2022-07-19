package minOperations

import (
	"reflect"
	"testing"
)

func Test_minOperations(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 []int
		expect int
	}{
		{"case 1", []int{2, 3, 2, 4, 3}, []int{9, 6, 9, 3, 15}, 2},
		{"case 2", []int{4, 3, 6}, []int{8, 2, 6, 10}, -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minOperations(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
