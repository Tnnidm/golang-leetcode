package findSmallestSetOfVertices

import (
	"reflect"
	"testing"
)

func Test_findSmallestSetOfVertices(t *testing.T) {
	cases := []struct {
		name   string
		input1 int
		input2 [][]int
		expect []int
	}{
		{"case 1", 5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}, []int{0, 2, 3}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findSmallestSetOfVertices(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
