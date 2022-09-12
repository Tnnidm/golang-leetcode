package partitionString

import (
	"reflect"
	"testing"
)

func Test_partitionString(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := partitionString(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
