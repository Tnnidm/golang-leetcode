package reorderList

import (
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
	cases := []struct {
		name   string
		input  *ListNode
		expect *ListNode
	}{
		{"case 1", &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, &ListNode{1, &ListNode{4, &ListNode{2, &ListNode{3, nil}}}}},
		{"case 1", &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, &ListNode{1, &ListNode{5, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			reorderList(c.input)
			if !reflect.DeepEqual(c.input, c.expect) {
				t.Fatalf("expect: %v, but got: %v", c.expect, c.input)
			}
		})
	}
}
