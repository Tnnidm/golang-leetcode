package minNumber

import (
	"reflect"
	"testing"
)

func TestMinNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect string
	}{
		{"case 1", []int{10, 2}, "102"},
		{"case 2", []int{3, 30, 34, 5, 9}, "3033459"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minNumber(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
