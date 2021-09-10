package removeElement

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		nums   []int
		val    int
		expect int
	}{
		{"1 test 1", []int{3, 2, 2, 3}, 3, 2},
		{"2 test 2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := removeElement(c.nums, c.val)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, ret, c.nums, c.val)
			}
		})
	}
}
