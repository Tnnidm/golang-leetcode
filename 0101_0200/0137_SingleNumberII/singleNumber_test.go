package singleNumber

import (
	"reflect"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{4, -1, 4, 4}, -1},
		{"case 2", []int{9, 1, 7, 9, 7, 9, 7}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := singleNumber(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
