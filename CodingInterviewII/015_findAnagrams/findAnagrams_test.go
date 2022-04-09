package findAnagrams

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect []int
	}{
		{"case 1", "cbaebabacd", "abc", []int{0, 6}},
		{"case 2", "abab", "ab", []int{0, 1, 2}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findAnagrams(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
