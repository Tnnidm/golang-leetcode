package lastRemaining

import (
	"reflect"
	"testing"
)

func Test_lastRemaining(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		m      int
		expect int
	}{
		{"case 1", 5, 3, 3},
		{"case 2", 10, 17, 2},
		{"case 3", 5, 1, 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := lastRemaining(c.n, c.m)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.n, c.m)
			}
		})
	}
}
