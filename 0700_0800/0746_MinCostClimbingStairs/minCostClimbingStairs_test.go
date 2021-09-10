package minCostClimbingStairs

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
		{"1 test 1", []int{10, 15, 20}, 15},
		{"2 test 2", []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minCostClimbingStairs(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
