package singleNumbers

import (
	"reflect"
	"testing"
)

func Test_twoSingleNumbers(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{4, 1, 4, 6}, []int{1, 6}},
		{"case 2", []int{0, 2}, []int{0, 2}},
		{"case 3", []int{1, 2, 10, 4, 1, 4, 3, 3}, []int{2, 10}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := twoSingleNumbers(c.input)
			if !reflect.DeepEqual(ret, c.expect) && !reflect.DeepEqual(ret, exchange(c.expect)) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}

func exchange(ans []int) []int {
	ans[0], ans[1] = ans[1], ans[0]
	return ans
}

func Test_singleNumberInTriple(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{4, 1, 4, 4}, 1},
		{"case 2", []int{9, 1, 7, 9, 7, 9, 7}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := singleNumberInTriple(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
