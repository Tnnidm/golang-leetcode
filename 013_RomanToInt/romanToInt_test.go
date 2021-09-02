package romanToInt

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"1 test 1", "III", 3},
		{"2 test 2", "IV", 4},
		{"3 test 3", "IX", 9},
		{"4 test 4", "LVIII", 58},
		{"5 test 5", "MCMXCIV", 1994},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := romanToInt(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
