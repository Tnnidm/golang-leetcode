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
		{"case 1", []int{2, 7, 11, 15}, 9, []int{2, 7}},
		{"case 2", []int{10, 26, 30, 31, 47, 60}, 40, []int{10, 30}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := twoSum(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) && !reflect.DeepEqual(ret, exchange(c.expect)) {
				t.Fatalf("expect: %v but got: %v with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}

func exchange(ans []int) []int {
	ans[0], ans[1] = ans[1], ans[0]
	return ans
}
