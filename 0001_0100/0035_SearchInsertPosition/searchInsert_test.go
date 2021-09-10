package searchInsert

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		nums   []int
		target int
		expect int
	}{
		{"1 test 1", []int{1, 3, 5, 6}, 5, 2},
		{"2 test 2", []int{1, 3, 5, 6}, 2, 1},
		{"3 test 3", []int{1, 3, 5, 6}, 7, 4},
		{"4 test 4", []int{1, 3, 5, 6}, 0, 0},
		{"5 test 5", []int{1}, 0, 0},
		{"6 test 6", []int{1, 3}, 1, 0},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := searchInsert(c.nums, c.target)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, ret, c.nums, c.target)
			}
		})
	}
}
