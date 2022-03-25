package intToRoman

import (
	"reflect"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect string
	}{
		{"case 1", 1, "I"},
		{"case 2", 3, "III"},
		{"case 3", 4, "IV"},
		{"case 4", 7, "VII"},
		{"case 5", 58, "LVIII"},
		{"case 6", 1994, "MCMXCIV"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := intToRoman(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
