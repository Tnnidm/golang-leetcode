package minCut

import (
	"reflect"
	"testing"
)

func Test_minCut(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"case 1", "aab", 1},
		{"case 2", "a", 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minCut(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
