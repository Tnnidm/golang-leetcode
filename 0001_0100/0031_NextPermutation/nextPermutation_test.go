package nextPermutation

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"case 2", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case 3", []int{1, 1, 5}, []int{1, 5, 1}},
	}
	for i, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			nextPermutation(c.input)
			if !reflect.DeepEqual(c.input, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, c.input, cases[i].input)
			}
		})
	}
}
