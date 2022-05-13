package minFlipsMonoIncr

import (
	"reflect"
	"testing"
)

func Test_minFlipsMonoIncr(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"case 1", "00110", 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minFlipsMonoIncr(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
