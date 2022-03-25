package fourSum

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect [][]int
	}{
		{"case 1", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"case 2", []int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := fourSum(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
