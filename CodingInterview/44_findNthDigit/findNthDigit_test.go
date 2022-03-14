package findNthDigit

import (
	"reflect"
	"testing"
)

func TestFindNthDigit(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{"case 1", 3, 3},
		{"case 2", 11, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findNthDigit(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
