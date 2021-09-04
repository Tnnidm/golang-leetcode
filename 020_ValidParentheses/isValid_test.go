package isValid

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  string
		expect bool
	}{
		{"1 test 1", "()", true},
		{"2 test 2", "()[]{}", true},
		{"3 test 3", "(]", false},
		{"4 test 4", "([)]", false},
		{"5 test 5", "{[]}", true},
		{"6 test 6", "{[]})", false},
		{"7 test 7", "){[]}", false},
		{"8 test 8", "{", false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isValid(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
