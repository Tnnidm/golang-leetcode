package printNumbers

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect []string
	}{
		{"case 1", 1, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := printNumbers(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
