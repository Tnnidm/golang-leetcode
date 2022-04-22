package countAndSay

import (
	"reflect"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect string
	}{
		{"case 1", 1, "1"},
		{"case 1", 4, "1211"},
		{"case 1", 5, "111221"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := countAndSay(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
