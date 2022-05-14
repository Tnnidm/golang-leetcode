package longestCommonSubsequence

import (
	"reflect"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect int
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestCommonSubsequence(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
