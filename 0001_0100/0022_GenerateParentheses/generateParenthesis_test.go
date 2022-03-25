package generateParenthesis

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect []string
	}{
		{"case 0", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"case 1", 1, []string{"()"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := generateParenthesis(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
