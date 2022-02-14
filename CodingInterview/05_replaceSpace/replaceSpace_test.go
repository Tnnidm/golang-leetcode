package replaceSpace

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"case 1", "ab cd", "ab%20cd"},
		{"case 2", "abcd", "abcd"},
		{"case 3", "", ""},
		{"case 4", " ", "%20"},
		{"case 5", "ab   cd", "ab%20%20%20cd"},
		{"case 6", "    ", "%20%20%20%20"},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := replaceSpace(c.input)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
				}
			},
		)
	}
}
