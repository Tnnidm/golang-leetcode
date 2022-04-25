package findMaximumXOR

import (
	"reflect"
	"testing"
)

func Test_findMaximumXOR(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{3, 10, 5, 25, 2, 8}, 28},
		{"case 2", []int{0}, 0},
		{"case 3", []int{2, 4}, 6},
		{"case 4", []int{8, 10, 2}, 10},
		{"case 5", []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}, 127},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findMaximumXOR(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
