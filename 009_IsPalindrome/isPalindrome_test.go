package isPalindrome

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  int
		expect bool
	}{
		{"1 test 1", 121, true},
		{"2 test 2", -121, false},
		{"3 test 3", 10, false},
		{"4 test 4", -101, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isPalindrome(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
