package simplifyPath

import (
	"reflect"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"case 1", "/home//foo/", "/home/foo"},
		{"case 2", "/home/../foo/", "/foo"},
		{"case 3", "/home/.././foo/", "/foo"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := simplifyPath(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
