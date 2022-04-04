package removeComments

import (
	"reflect"
	"testing"
)

func Test_removeComments(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect []string
	}{
		{"case 1", []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}, []string{"int main()", "{ ", "  ", "int a, b, c;", "a = b + c;", "}"}},
		{"case 2", []string{"a/*comment", "line", "more_comment*/b"}, []string{"ab"}},
		{"case 3", []string{"a///*comment", "line", "more_comment*/b"}, []string{"line", "more_comment*/b"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := removeComments(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
