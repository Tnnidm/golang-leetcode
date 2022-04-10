package countSubstrings

import (
	"reflect"
	"testing"
)

func Test_countSubstrings(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"case 1", "abc", 3},
		{"case 2", "aaa", 6},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := countSubstrings(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
