package findKthLargest

import (
	"reflect"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect int
	}{
		{"case 1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findKthLargest(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
