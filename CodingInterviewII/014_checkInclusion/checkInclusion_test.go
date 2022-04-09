package checkInclusion

import (
	"reflect"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect bool
	}{
		{"case 1", "ab", "eidbaooo", true},
		{"case 2", "ab", "eidboaoo", false},
		{"case 3", "abcd", "abc", false},
		{"case 4", "abcd", "abce", false},
		{"case 5", "abcd", "abdc", true},
		{"case 6", "abcd", "abc", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := checkInclusion(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
