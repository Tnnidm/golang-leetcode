package minWindow

import (
	"reflect"
	"testing"
)

func Test_minWindow(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect string
	}{
		{"case 1", "ADOBECODEBANC", "ABC", "BANC"},
		{"case 2", "a", "a", "a"},
		{"case 3", "a", "aa", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minWindow(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
