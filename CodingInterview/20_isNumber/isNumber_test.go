package isNumber

import (
	"reflect"
	"testing"
)

func Test_isNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect bool
	}{
		{"case 1", "+1", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isNumber(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
