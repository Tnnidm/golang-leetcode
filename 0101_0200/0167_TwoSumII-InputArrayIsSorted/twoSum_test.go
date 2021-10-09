package twoSum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		numbers []int
		target  int
		expect  []int
	}{
		{"case 1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"case 2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"case 3", []int{-1, 0}, -1, []int{1, 2}},
		// {"2 test 2", "race a car", false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := twoSum(c.numbers, c.target)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, ret, c.numbers, c.target)
			}
		})
	}
}
