package minGroups

import (
	"reflect"
	"testing"
)

func Test_minGroups(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{"case 1", [][]int{{5, 10}, {6, 9}, {1, 5}, {2, 3}, {1, 10}}, 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minGroups(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
