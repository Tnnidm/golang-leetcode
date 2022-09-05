package isStrictlyPalindromic

import (
	"reflect"
	"testing"
)

func Test_isStrictlyPalindromic(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect bool
	}{
		{"case 1", 9, false},
		{"case 2", 4, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isStrictlyPalindromic(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
