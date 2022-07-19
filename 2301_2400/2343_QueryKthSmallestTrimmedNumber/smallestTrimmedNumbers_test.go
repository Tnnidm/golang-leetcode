package smallestTrimmedNumbers

import (
	"reflect"
	"testing"
)

func Test_smallestTrimmedNumbers(t *testing.T) {
	cases := []struct {
		name   string
		input1 []string
		input2 [][]int
		expect []int
	}{
		{"case 1", []string{"102", "473", "251", "814"}, [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}}, []int{2, 2, 1, 0}},
		{"case 2", []string{"24", "37", "96", "04"}, [][]int{{2, 1}, {2, 2}}, []int{3, 0}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := smallestTrimmedNumbers(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
