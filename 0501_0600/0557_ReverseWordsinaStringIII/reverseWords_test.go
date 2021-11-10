package reverseWords

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"case 1", "Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"case 2", "God Ding", "doG gniD"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverseWords(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
