package longestContinuousSubstring

import (
	"reflect"
	"testing"
)

func Test_longestContinuousSubstring(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"case 1", "abacaba", 2},
		{"case 2", "abcde", 5},
		{"case 3", "a", 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestContinuousSubstring(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
