package threeSumClosest

import (
	"reflect"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect int
	}{
		{"case 1", []int{-1, 2, 1, -4}, 1, 2},
		{"case 2", []int{0, 0, 0}, 1, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := threeSumClosest(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
