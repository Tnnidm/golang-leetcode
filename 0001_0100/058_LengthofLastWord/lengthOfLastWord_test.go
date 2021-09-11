package lengthOfLastWord

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
		{"1 test 1", "Hello World", 5},
		{"2 test 2", "   fly me   to   the moon  ", 4},
		{"3 test 3", "luffy is still joyboy", 6},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := lengthOfLastWord(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
