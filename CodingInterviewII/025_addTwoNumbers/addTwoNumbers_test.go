package addTwoNumbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	cases := []struct {
		name   string
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{"case 1", &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}, &ListNode{7, &ListNode{8, &ListNode{0, &ListNode{7, nil}}}}},
		{"case 2", &ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}, &ListNode{8, &ListNode{0, &ListNode{7, nil}}}},
		{"case 3", &ListNode{0, nil}, &ListNode{0, nil}, &ListNode{0, nil}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addTwoNumbers(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
