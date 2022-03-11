package verifyPostorder

import (
	"reflect"
	"testing"
)

func TestVerifyPostorder(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect bool
	}{
		{"case 1", []int{}, true},
		{"case 2", []int{2}, true},
		{"case 3", []int{1, 2}, true},
		{"case 4", []int{3, 2}, true},
		{"case 5", []int{3, 1, 2}, false},
		{"case 6", []int{1, 3, 2}, true},
		{"case 7", []int{1, 6, 3, 2, 5}, false},
		{"case 8", []int{1, 3, 2, 6, 5}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := verifyPostorder(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
