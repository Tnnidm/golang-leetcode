package isAnagram

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		s      string
		t      string
		expect bool
	}{
		{"case 1", "anagram", "nagaram", true},
		{"case 2", "rat", "cat", false},
		{"case 3", "nl", "cx", false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isAnagram(c.s, c.t)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, ret, c.s, c.t)
			}
		})
	}
}
