package minNumberOfHours

import (
	"reflect"
	"testing"
)

func Test_minNumberOfHours(t *testing.T) {
	cases := []struct {
		name   string
		input1 int
		input2 int
		input3 []int
		input4 []int
		expect int
	}{
		{"case 1", 5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1}, 8},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minNumberOfHours(c.input1, c.input2, c.input3, c.input4)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v", c.expect, ret)
			}
		})
	}
}
