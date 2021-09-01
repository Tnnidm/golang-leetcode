package reverse

import (
	"math"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{"1 test 1", 123, 321},
		{"2 test 2", -123, -321},
		{"3 test 3", 1234000, 4321},
		{"4 test 4", int(math.Pow(2, 31) - 1), 0},
		{"5 test 5", 0, 0},
		{"5 test 5", 1534236469, 0},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverse(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
