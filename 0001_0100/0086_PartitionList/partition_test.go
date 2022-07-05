package partition

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	cases := []struct {
		name   string
		input1 *ListNode
		input2 int
		expect *ListNode
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := partition(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
