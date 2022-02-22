package hammingWeight

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input  uint32
		expect int
	}{
		{"case 1", 11, 3},
		{"case 2", 128, 1},
		{"case 3", 4294967293, 31},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := hammingWeight(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
