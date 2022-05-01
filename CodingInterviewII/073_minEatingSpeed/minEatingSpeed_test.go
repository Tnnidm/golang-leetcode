package minEatingSpeed

import (
	"reflect"
	"testing"
)

func Test_minEatingSpeed(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect int
	}{
		{"case 1", []int{3, 6, 7, 11}, 8, 4},
		{"case 2", []int{30, 11, 23, 4, 20}, 5, 30},
		{"case 3", []int{30, 11, 23, 4, 20}, 6, 23},
		{"case 3", []int{30, 11, 23}, 64, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minEatingSpeed(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
