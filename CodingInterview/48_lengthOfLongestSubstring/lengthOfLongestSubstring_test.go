package lengthOfLongestSubstring

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"case 1", "abcabcbb", 3},
		{"case 2", "bbbbb", 1},
		{"case 3", "pwwkew", 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := lengthOfLongestSubstring(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
