package findSubarrays

import (
	"reflect"
	"testing"
)

func Test_findSubarrays(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect bool
	}{
		{"case 1", []int{77, 95, 90, 98, 8, 100, 88, 96, 6, 40, 86, 56, 98, 96, 40, 52, 30, 33, 97, 72, 54, 15, 33, 77, 78, 8, 21, 47, 99, 48}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findSubarrays(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
