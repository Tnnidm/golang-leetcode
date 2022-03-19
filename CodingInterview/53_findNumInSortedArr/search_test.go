package findNumInSortedArr

import (
	"reflect"
	"testing"
)

func TestReversePairs(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect int
	}{
		{"case 1", []int{5, 7, 7, 8, 8, 10}, 8, 2},
		{"case 2", []int{5, 7, 7, 8, 8, 10}, 6, 0},
		{"case 3", []int{1, 1, 1, 1}, 1, 4},
		{"case 4", []int{1}, 1, 1},
		{"case 5", []int{2, 2}, 1, 0},
		{"case 6", []int{2, 2}, 3, 0},
		{"case 7", []int{}, 1, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := search(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
