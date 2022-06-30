package uniquePathsWithObstacles

import (
	"reflect"
	"testing"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{"case 1", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := uniquePathsWithObstacles(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
