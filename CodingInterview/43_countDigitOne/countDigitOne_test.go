package countDigitOne

import (
	"reflect"
	"testing"
)

func TestCountDigitOne(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{"case 1", 1, 1},
		{"case 2", 5, 1},
		{"case 3", 10, 2},
		{"case 4", 12, 5},
		{"case 4", 22, 13},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := countDigitOne(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v. with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
