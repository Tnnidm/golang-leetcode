package addTwoNumbers

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{"case 1", &ListNode{9, nil}, &ListNode{1, nil}, &ListNode{0, &ListNode{1, nil}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addTwoNumbers(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v and %v",
					c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
