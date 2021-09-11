package maxProfit

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
		{"1 test 1", []int{7, 1, 5, 3, 6, 4}, 5},
		// {"2 test 2", []int{7, 6, 4, 3, 1}, 0},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxProfit(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
