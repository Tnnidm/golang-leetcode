package minimumLengthEncoding

import (
	"reflect"
	"testing"
)

func Test_minimumLengthEncoding(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect int
	}{
		{"case 1", []string{"time", "me", "bell"}, 10},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minimumLengthEncoding(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
