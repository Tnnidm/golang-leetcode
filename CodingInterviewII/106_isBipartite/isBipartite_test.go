package isBipartite

import (
	"reflect"
	"testing"
)

func Test_isBipartite(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect bool
	}{
		{"case 1", [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, false},
		{"case 2", [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isBipartite(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
