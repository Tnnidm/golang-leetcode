package checkDistances

import (
	"reflect"
	"testing"
)

func Test_checkDistances(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 []int
		expect bool
	}{
		{"case 1", "abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, true},
		{"case 2", "aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := checkDistances(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v, %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
