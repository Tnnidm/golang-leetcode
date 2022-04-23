package containsNearbyAlmostDuplicate

import (
	"reflect"
	"testing"
)

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		input3 int
		expect bool
	}{
		{"case 1", []int{1, 2, 3, 1}, 3, 0, true},
		{"case 1", []int{1, 0, 1, 1}, 1, 2, true},
		{"case 1", []int{1, 5, 9, 1, 5, 9}, 2, 3, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := containsNearbyAlmostDuplicate(c.input1, c.input2, c.input3)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v, %v and %v", c.expect, ret, c.input1, c.input2, c.input3)
			}
		})
	}
}
