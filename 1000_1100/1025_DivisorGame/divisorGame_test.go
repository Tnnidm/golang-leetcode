package divisorGame

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  int
		expect bool
	}{
		{"1 test 1", 2, true},
		{"2 test 2", 3, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := divisorGame(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
