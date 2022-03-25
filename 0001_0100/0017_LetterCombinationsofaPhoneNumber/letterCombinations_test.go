package letterCombinations

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect []string
	}{
		{"case 1", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		// {"case 2", "", []string{}},
		// {"case 3", "2", []string{"a", "b", "c"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := letterCombinations(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
