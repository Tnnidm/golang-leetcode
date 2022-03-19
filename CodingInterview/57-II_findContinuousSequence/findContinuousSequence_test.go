package findContinuousSequence

import (
	"reflect"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect [][]int
	}{
		{"case 1", 9, [][]int{{2, 3, 4}, {4, 5}}},
		// {"case 2", 15, [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findContinuousSequence(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
