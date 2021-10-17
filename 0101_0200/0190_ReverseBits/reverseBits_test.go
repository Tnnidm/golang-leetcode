package reverseBits

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  uint32
		expect uint32
	}{
		{"case 1", 43261596, 964176192},
		{"case 2", 4294967293, 3221225471},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverseBits(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
