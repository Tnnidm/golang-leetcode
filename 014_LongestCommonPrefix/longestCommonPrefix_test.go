package longestCommonPrefix

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  []string
		expect string
	}{
		{"1 test 1", []string{"flower", "flow", "flight"}, "fl"},
		{"2 test 2", []string{"dog", "racecar", "car"}, ""},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestCommonPrefix(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
