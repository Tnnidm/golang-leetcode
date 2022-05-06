package combinationSum

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect [][]int
	}{
		{"case 1", []int{2, 3, 6, 7}, 7, [][]int{{7}, {2, 2, 3}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := combinationSum(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
