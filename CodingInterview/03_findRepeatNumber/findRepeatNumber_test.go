package findRepeatNumber

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"test case 1", []int{2, 3, 1, 0, 2, 5, 3}, 2},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := findRepeatNumber(c.input)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expected: %v, but got: %v, with input: %v", c.expect, ret, c.input)
				}
			},
		)
	}
}
