package dailyTemperatures

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"case 2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"case 3", []int{30, 60, 90}, []int{30, 60, 90}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := dailyTemperatures(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
