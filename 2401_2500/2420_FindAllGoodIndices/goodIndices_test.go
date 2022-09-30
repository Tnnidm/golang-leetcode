package goodIndices

import (
	"reflect"
	"testing"
)

func Test_goodIndices(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect []int
	}{
		{"case 1", []int{2, 1, 1, 1, 3, 4, 1}, 2, []int{2, 3}},
		{"case 2", []int{2, 1, 1, 2}, 2, []int{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := goodIndices(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
