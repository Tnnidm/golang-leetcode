package combinationSum2

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect [][]int
	}{
		// {"case 1", []int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{"case 1", []int{1, 2, 3}, 3, [][]int{{1, 2}, {3}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := combinationSum2(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
