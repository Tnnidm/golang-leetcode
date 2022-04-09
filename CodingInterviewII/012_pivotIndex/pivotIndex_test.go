package pivotIndex

import (
	"reflect"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		expect int
	}{
		{"case 1", []int{1, 7, 3, 6, 5, 6}, 3},
		{"case 1", []int{1, 2, 3}, -1},
		{"case 1", []int{2, 1, -1}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := pivotIndex(c.input1)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input1)
			}
		})
	}
}
