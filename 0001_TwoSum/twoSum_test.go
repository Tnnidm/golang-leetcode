package twoSum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect []int
	}{
		{"1 test 1", [][]int{{2, 7, 11, 15}, {9}}, []int{0, 1}},
		{"2 test 2", [][]int{{3, 2, 4}, {6}}, []int{1, 2}},
		{"3 test 3", [][]int{{2, 7, 11, 15}, {9}}, []int{0, 1}},
		{"4 test 4", [][]int{{7, 6, 5, 3, 2, 1, 4, 9, 10}, {17}}, []int{0, 8}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := twoSum(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
