package singleNonDuplicate

import (
	"reflect"
	"testing"
)

func Test_singleNonDuplicate(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		// {"case 1", []int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		// {"case 2", []int{3, 3, 7, 7, 10, 11, 11}, 10},
		// {"case 3", []int{1, 3, 3, 7, 7, 11, 11}, 1},
		{"case 4", []int{3, 3, 7, 7, 11, 11, 12}, 12},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := singleNonDuplicate(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
