package PascalsTriangleII

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  int
		expect []int
	}{
		{"1 test 1", 3, []int{1, 3, 3, 1}},
		{"3 test 3", 0, []int{1}},
		{"4 test 4", 1, []int{1, 1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := getRow(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
