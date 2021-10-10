package convertToTitle

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  int
		expect string
	}{
		{"case 1", 1, "A"},
		{"case 2", 28, "AB"},
		{"case 3", 701, "ZY"},
		{"case 4", 2147483647, "FXSHRXW"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := convertToTitle(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
