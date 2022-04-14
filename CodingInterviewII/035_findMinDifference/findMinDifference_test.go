package findMinDifference

import (
	"reflect"
	"testing"
)

func Test_findMinDifference(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect int
	}{
		{"case 1", []string{"23:59", "00:00"}, 1},
		{"case 2", []string{"00:00", "23:59", "00:00"}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findMinDifference(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
