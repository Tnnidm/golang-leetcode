package reverseList

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	cases := []struct {
		name   string
		input  *ListNode
		expect *ListNode
	}{
		{"case 1", &ListNode{1, nil}, &ListNode{1, nil}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverseList(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
