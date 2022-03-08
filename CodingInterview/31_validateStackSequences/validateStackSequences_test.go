package validateStackSequences

import (
	"reflect"
	"testing"
)

func Test_validateStackSequences(t *testing.T) {
	cases := []struct {
		name   string
		pushed []int
		poped  []int
		expect bool
	}{
		{"case 1", []int{}, []int{}, true},
		{"case 2", []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{"case 3", []int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
		{"case 4", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, true},
		{"case 5", []int{1}, []int{1}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := validateStackSequences(c.pushed, c.poped)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.pushed, c.poped)
			}
		})
	}
}
