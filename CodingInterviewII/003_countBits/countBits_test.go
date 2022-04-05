package countBits

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect []int
	}{
		{"case 1", 0, []int{0}},
		{"case 2", 1, []int{0, 1}},
		{"case 3", 2, []int{0, 1, 1}},
		{"case 4", 5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := countBits(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
