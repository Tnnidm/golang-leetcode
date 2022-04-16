package maximalRectangle

import (
	"reflect"
	"testing"
)

func Test_maximalRectangle(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect int
	}{
		{"case 1", []string{"10100", "10111", "11111", "10010"}, 6},
		{"case 2", []string{}, 0},
		{"case 3", []string{"0"}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maximalRectangle(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
