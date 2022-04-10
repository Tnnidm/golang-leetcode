package removeNthFromEnd

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	cases := []struct {
		name   string
		input1 *ListNode
		input2 int
		expect *ListNode
	}{
		{"case 1", &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}}},
		{"case 2", &ListNode{1, nil}, 1, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := removeNthFromEnd(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
