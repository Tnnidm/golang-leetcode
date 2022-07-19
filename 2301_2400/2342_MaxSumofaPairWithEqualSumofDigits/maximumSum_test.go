package maximumSum

import (
	"reflect"
	"testing"
)

func Test_maximumSum(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{18, 43, 36, 13, 7}, 54},
		{"case 2", []int{10, 12, 19, 14}, -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maximumSum(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
