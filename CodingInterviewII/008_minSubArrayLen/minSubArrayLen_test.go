package minSubArrayLen

import (
	"reflect"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	cases := []struct {
		name   string
		input1 int
		input2 []int
		expect int
	}{
		{"case 1", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"case 2", 4, []int{1, 4, 4}, 1},
		{"case 3", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minSubArrayLen(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
