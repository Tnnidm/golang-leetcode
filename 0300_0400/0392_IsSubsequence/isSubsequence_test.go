package isSubsequence

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect bool
	}{
		{"1 test 1", "abc", "ahbgdc", true},
		{"2 test 2", "axc", "ahbgdc", false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isSubsequence(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
