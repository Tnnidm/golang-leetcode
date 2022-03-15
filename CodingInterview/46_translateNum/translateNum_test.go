package translateNum

import (
	"reflect"
	"testing"
)

func TestMinNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{"case 1", 11, 2},
		{"case 2", 12258, 5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := translateNum(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
