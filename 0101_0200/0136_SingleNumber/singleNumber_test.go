package singleNumber

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
		{"case 1", []int{2, 2, 1}, 2},
		{"case 2", []int{4, 1, 2, 1, 2}, 4},
		{"case 3", []int{1}, 1},
		// {"2 test 2", "race a car", false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := singleNumber(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
