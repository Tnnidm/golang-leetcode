package isStraight

import (
	"reflect"
	"testing"
)

func Test_dicesProbability(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect bool
	}{
		{"case 1", []int{1, 2, 3, 4, 5}, true},
		{"case 2", []int{1, 2, 4, 4, 5}, false},
		{"case 3", []int{0, 0, 1, 2, 5}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isStraight(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
