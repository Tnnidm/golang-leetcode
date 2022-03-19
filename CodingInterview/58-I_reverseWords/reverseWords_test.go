package reverseWords

import (
	"reflect"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"case 1", "the sky is blue", "blue is sky the"},
		{"case 2", "  hello world!  ", "world! hello"},
		{"case 3", "a good   example", "example good a"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverseWords(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
