package intersect

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 []int
		input2 []int
		expect []int
	}{
		{"case 1", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := intersect(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
