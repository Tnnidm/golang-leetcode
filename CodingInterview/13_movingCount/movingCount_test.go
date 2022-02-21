package movingCount

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input1 int
		input2 int
		input3 int
		expect int
	}{
		{"case 1", 2, 3, 1, 3},
		{"case 2", 3, 1, 0, 1},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := movingCount(c.input1, c.input2, c.input3)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v , %v and %v", c.expect, ret, c.input1, c.input2, c.input3)
				}
			},
		)
	}
}
