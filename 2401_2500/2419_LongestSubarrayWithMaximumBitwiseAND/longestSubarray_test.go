package longestSubarray

import (
	"reflect"
	"testing"
)

func Test_longestSubarray(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{1, 2, 3, 3, 2, 2}, 2},
		{"case 2", []int{1, 2, 3, 4}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestSubarray(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
