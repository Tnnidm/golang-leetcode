package reorderArray

import (
	"reflect"
	"testing"
)

func Test_isNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := exchange(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
