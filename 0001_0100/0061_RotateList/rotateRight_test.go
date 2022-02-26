package rotateRight

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 *ListNode
		input2 int
		expect *ListNode
	}{
		{"case 1", &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 1, &ListNode{4, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := rotateRight(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v and %v",
					c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
