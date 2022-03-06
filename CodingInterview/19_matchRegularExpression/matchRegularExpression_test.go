package matchRegularExpression

import (
	"reflect"
	"testing"
)

func TestMatchRegularExpression(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect bool
	}{
		{"case 1", "aa", "a", false},
		{"case 2", "aa", "a*", true},
		{"case 3", "ab", ".*", true},
		{"case 4", "aaa", "a.a", true},
		{"case 5", "aaa", "ab*ac*a", true},
		{"case 6", "aaa", "aa.a", false},
		{"case 7", "aaa", "ab*a", false},
		{"case 8", "a", "ab*", true},
		{"case 9", "abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*b*", true},
		{"case 10", "c", ".*b*b*", true},
		{"case 11", "aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isMatch(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
