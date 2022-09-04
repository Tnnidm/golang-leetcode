package largestPalindromic

import (
	"reflect"
	"testing"
)

func Test_largestPalindromic(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"case 1", "444947137", "7449447"},
		{"case 2", "00009", "9"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := largestPalindromic(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
