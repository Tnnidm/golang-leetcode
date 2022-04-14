package isAnagram

import (
	"reflect"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect bool
	}{
		{"case 1", "anagram", "nagaram", true},
		{"case 2", "rat", "car", false},
		{"case 3", "a", "a", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isAnagram(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
