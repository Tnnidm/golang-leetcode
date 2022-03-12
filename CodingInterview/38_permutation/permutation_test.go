package permutation

import (
	"reflect"
	"testing"
)

func TestTreeToDoublyList(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect []string
	}{
		{"case 1", "abc", []string{"abc", "acb", "bac", "bca", "cba", "cab"}},
		{"case 2", "ab", []string{"ab", "ba"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := permutation(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
