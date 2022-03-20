package constructArr

import (
	"reflect"
	"testing"
)

func Test_constructArr(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := constructArr(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
