package restoreIpAddresses

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect []string
	}{
		{"case 1", "25525511135", []string{"255.255.11.135", "255.255.111.35"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := restoreIpAddresses(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
