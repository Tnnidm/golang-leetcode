package validPalindrome

import (
	"reflect"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect bool
	}{
		{"case 1", "aba", true},
		{"case 2", "abca", true},
		{"case 3", "abc", false},
		{"case 4", "a", true},
		{"case 4", "ab", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := validPalindrome(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
