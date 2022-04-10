package isPalindrome

import (
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
	cases := []struct {
		name   string
		input  *ListNode
		expect bool
	}{
		{"case 1", &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isPalindrome(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
