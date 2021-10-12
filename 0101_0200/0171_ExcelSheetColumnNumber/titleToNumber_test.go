package titleToNumber

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
		{"case 1", "A", 1},
		{"case 2", "AB", 28},
		{"case 3", "ZY", 701},
		{"case 4", "FXSHRXW", 2147483647},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := titleToNumber(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
