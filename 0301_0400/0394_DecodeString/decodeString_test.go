package decodeString

import (
	"reflect"
	"testing"
)

func Test_decodeString(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := decodeString(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
