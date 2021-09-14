package addBinary

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
		expect string
	}{
		{"1 test 1", "11", "1", "100"},
		{"2 test 2", "1010", "1011", "10101"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addBinary(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with input1: %v, input2: %v",
					c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
