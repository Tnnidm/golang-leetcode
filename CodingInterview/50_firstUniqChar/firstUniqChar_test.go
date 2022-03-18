package firstUniqChar

import (
	"reflect"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect byte
	}{
		{"case 1", "abaccdeff", 'b'},
		{"case 2", "ababcdcdefef", ' '},
		{"case 3", "", ' '},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := firstUniqChar(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
