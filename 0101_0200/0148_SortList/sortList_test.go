package sortList

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	cases := []struct {
		name   string
		input  *ListNode
		expect *ListNode
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := sortList(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
