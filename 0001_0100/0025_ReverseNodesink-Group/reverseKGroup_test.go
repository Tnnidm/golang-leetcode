package reverseKGroup

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	cases := []struct {
		name   string
		input1 *ListNode
		input2 int
		expect int
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverseKGroup(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
