package convert

import (
	"reflect"
	"testing"
)

func Test_convert(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 int
		expect string
	}{
		{"case 1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"case 2", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"case 3", "A", 1, "A"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := convert(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
