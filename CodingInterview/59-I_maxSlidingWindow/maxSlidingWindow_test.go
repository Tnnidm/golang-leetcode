package maxSlidingWindow

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect []int
	}{
		{"case 1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"case 2", []int{1, 3, -1, -3, 5, 3, 6, 7}, 1, []int{1, 3, -1, -3, 5, 3, 6, 7}},
		{"case 3", []int{1, 3, -1, -3, 5, 3, 6, 7}, 7, []int{6, 7}},
		{"case 4", []int{7, 6, 5, 4, 3, 2, 1, 0}, 3, []int{7, 6, 5, 4, 3, 2}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxSlidingWindow(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
