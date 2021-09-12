package plusOne

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"1 test 1", []int{1, 1, 2}, []int{1, 1, 3}},
		{"2 test 2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"3 test 3", []int{9}, []int{1, 0}},
		{"4 test 4", []int{9, 9, 9, 9}, []int{1, 0, 0, 0, 0}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := plusOne(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
