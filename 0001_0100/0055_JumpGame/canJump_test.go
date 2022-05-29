package canJump

import (
	"reflect"
	"testing"
)

func Test_canJump(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect bool
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := canJump(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
