package deleteDuplicates

import (
	"reflect"
	"testing"
)

func transfer(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
func Test(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input  *ListNode
		expect *ListNode
	}{
		{"case 1", &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}, &ListNode{1, &ListNode{2, &ListNode{5, nil}}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := deleteDuplicates(c.input)
			if !reflect.DeepEqual(transfer(ret), transfer(c.expect)) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input)
			}
		})
	}
}
