package merge

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{"case 1", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"case 2", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := merge(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
