package maxSubArray

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"1 test 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"2 test 2", []int{1}, 1},
		{"3 test 3", []int{5, 4, -1, 7, 8}, 23},
		{"4 test 4", []int{-1}, -1},
		{"5 test 5", []int{-100000}, -100000},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxSubArray(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
