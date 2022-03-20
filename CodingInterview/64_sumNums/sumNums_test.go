package sumNums

import (
	"reflect"
	"testing"
)

func Test_sumNums(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{"case 1", 3, 6},
		{"case 2", 9, 45},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := sumNums(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
