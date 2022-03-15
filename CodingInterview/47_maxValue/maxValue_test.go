package maxValue

import (
	"reflect"
	"testing"
)

func TestMinNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{"case 1", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 12},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxValue(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
