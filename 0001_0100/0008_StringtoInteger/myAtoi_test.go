package myAtoi

import (
	"reflect"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"case 1", "-42", -42},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := myAtoi(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
