package findNumInSortedArr

import (
	"reflect"
	"testing"
)

func Test_getNumberSameAsIndex(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{-3, -1, 1, 3, 5}, 3},
		{"case 2", []int{0}, 0},
		{"case 2", []int{-1, 0, 2}, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := getNumberSameAsIndex(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
