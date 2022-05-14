package isInterleave

import (
	"reflect"
	"testing"
)

func Test_isInterleave(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		input3 string
		expect bool
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isInterleave(c.input1, c.input2, c.input3)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v, %v and %v", c.expect, ret, c.input1, c.input2, c.input3)
			}
		})
	}
}
