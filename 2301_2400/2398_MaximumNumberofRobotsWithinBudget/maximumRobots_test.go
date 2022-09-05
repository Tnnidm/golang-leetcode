package maximumRobots

import (
	"reflect"
	"testing"
)

func Test_maximumRobots(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		input2 []int
		input3 int64
		expect int
	}{
		{"case 1", []int{3, 6, 1, 3, 4}, []int{2, 1, 3, 4, 5}, 25, 3},
		{"case 1", []int{11, 12, 19}, []int{10, 8, 7}, 19, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maximumRobots(c.input1, c.input2, c.input3)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v, %v and %v", c.expect, ret, c.input1, c.input2, c.input3)
			}
		})
	}
}
