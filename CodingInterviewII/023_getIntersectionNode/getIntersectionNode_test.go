package getIntersectionNode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	sametail := &ListNode{2, &ListNode{4, nil}}
	cases := []struct {
		name   string
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{"case 1", &ListNode{0, &ListNode{9, &ListNode{1, sametail}}}, &ListNode{3, sametail}, sametail},
		{"case 2", &ListNode{2, &ListNode{6, &ListNode{4, nil}}}, &ListNode{1, &ListNode{5, nil}}, nil},
		{"case 3", nil, &ListNode{1, &ListNode{5, nil}}, nil},
		{"case 4", nil, nil, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := getIntersectionNode(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
