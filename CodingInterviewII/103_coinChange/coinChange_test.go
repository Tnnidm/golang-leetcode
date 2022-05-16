package coinChange

import (
	"reflect"
	"testing"
)

func Test_coinChange(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect int
	}{
		{"case 1", []int{1, 2, 5}, 11, 3},
		{"case 2", []int{2}, 3, -1},
		{"case 3", []int{1}, 0, 0},
		{"case 4", []int{1}, 1, 1},
		{"case 5", []int{1}, 2, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := coinChange(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
