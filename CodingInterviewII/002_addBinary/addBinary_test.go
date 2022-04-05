package addBinary

import (
	"reflect"
	"testing"
)

func Test_addBinary(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect string
	}{
		{"case 1", "11", "10", "101"},
		{"case 2", "1010", "1011", "10101"},
		{"case 3", "11111111", "1", "100000000"},
		{"case 4", "101111", "10", "110001"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addBinary(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
