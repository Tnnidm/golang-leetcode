package numberOfPairs

import (
	"reflect"
	"testing"
)

func Test_numberOfPairs(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{1, 3, 2, 1, 3, 2, 2}, []int{3, 1}},
		{"case 2", []int{1, 1}, []int{1, 0}},
		{"case 3", []int{0}, []int{0, 1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := numberOfPairs(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
