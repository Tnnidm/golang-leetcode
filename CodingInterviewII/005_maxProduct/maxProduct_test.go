package maxProduct

import (
	"reflect"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect int
	}{
		{"case 1", []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}, 16},
		{"case 2", []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
		{"case 3", []string{"a", "aa", "aaa", "aaaa"}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxProduct(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
