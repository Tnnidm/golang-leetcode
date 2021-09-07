package strStr

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		haystack string
		needle   string
		expect   int
	}{
		{"1 test 1", "hello", "ll", 2},
		{"2 test 2", "aaaaaaa", "bba", -1},
		{"3 test 3", "", "", 0},
		{"4 test 4", "aaa", "", 0},
		{"5 test 5", "a", "a", 0},
		{"6 test 6", "mississippi", "issi", 1},
		{"7 test 7", "aabaaabaaac", "aabaaac", 4},
		{"8 test 8", "abcdababcdab", "abcdabd", -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := strStr(c.haystack, c.needle)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, ret, c.haystack, c.needle)
			}
		})
	}
}
