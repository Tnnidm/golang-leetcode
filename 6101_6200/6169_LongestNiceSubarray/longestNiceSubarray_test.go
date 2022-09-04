package longestNiceSubarray

import (
	"reflect"
	"testing"
)

func Test_longestNiceSubarray(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		// {"case 1", []int{1, 3, 8, 48, 10}, 3},
		// {"case 2", []int{3, 1, 5, 11, 13}, 1},
		{"case 3", []int{615598711, 264, 65552, 50331652, 1, 1048576, 16384, 544, 270532608}, 8},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestNiceSubarray(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
