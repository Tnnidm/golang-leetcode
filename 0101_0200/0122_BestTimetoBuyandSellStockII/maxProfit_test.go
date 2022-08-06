package maxProfit

import (
	"reflect"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"case 2", []int{1, 2, 3, 4, 5}, 4},
		{"case 3", []int{7, 6, 4, 3, 1}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxProfit(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
