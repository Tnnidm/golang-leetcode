package longestPalindrome

import (
	"reflect"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"case 1", "babad", "bab"},
		{"case 2", "cbbd", "bb"},
		{"case 3", "ccc", "ccc"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestPalindrome(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
