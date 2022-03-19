package findNumInSortedArr

import (
	"reflect"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{0, 1, 3}, 2},
		{"case 2", []int{0, 1, 2, 3, 4, 5, 6, 7, 9}, 8},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := missingNumber(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
