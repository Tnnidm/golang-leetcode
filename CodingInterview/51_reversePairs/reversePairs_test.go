package reversePairs

import (
	"reflect"
	"testing"
)

func TestReversePairs(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case 1", []int{7, 5, 6, 4}, 5},
		{"case 2", []int{1, 2, 1, 3, 5, 7, 1, 2}, 8},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reversePairs(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
