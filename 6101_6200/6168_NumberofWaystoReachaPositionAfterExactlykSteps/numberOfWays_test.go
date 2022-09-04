package numberOfWays

import (
	"reflect"
	"testing"
)

func Test_numberOfWays(t *testing.T) {
	cases := []struct {
		name   string
		input1 int
		input2 int
		input3 int
		expect int
	}{
		{"case 1", 1, 2, 3, 3},
		{"case 2", 2, 5, 10, 0},
		{"case 3", 1, 1000, 999, 1},
		{"case 4", 989, 1000, 99, 934081896},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := numberOfWays(c.input1, c.input2, c.input3)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v, %v and %v", c.expect, ret, c.input1, c.input2, c.input3)
			}
		})
	}
}
