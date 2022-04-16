package asteroidCollision

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{5, 10, -5}, []int{5, 10}},
		{"case 2", []int{8, -8}, []int{}},
		{"case 3", []int{10, 2, -5}, []int{10}},
		{"case 4", []int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := asteroidCollision(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
