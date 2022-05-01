package relativeSortArray

import (
	"reflect"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 []int
		expect []int
	}{
		{"case 1", []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := relativeSortArray(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
