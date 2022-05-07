package partition

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect [][]string
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := partition(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
