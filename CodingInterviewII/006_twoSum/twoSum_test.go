package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect []int
	}{
		{"case 1", []int{1, 2, 4, 6, 10}, 8, []int{1, 3}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := twoSum(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
