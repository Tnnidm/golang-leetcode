package findMaxLength

import (
	"reflect"
	"testing"
)

func Test_findMaxLength(t *testing.T) {
	cases := []struct {
		name   string
		input1 []int
		expect int
	}{
		{"case 1", []int{0, 1}, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findMaxLength(c.input1)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input1)
			}
		})
	}
}
