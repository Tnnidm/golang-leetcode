package reversePrint

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input  *ListNode
		expect []int
	}{
		{"case 1", &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, []int{4, 3, 2, 1}},
		{"case 2", &ListNode{1, nil}, []int{1}},
		{"case 3", nil, []int{}},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := reversePrint(c.input)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
				}
			},
		)
	}
}
