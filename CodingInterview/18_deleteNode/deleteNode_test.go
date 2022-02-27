package deleteNode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	example1 := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	target1 := example1
	target2 := example1.Next
	target3 := example1.Next.Next.Next
	example2 := &ListNode{4, nil}
	cases := []struct {
		name   string
		input1 *ListNode
		input2 *ListNode
	}{
		{"case 1", example1, target1},
		{"case 2", example1, target2},
		{"case 3", example1, target3},
		{"case 4", example2, example2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := deleteNode(c.input1, c.input2)
			fmt.Println(c.name)
			for ; ret != nil; ret = ret.Next {
				fmt.Println(ret.Val)
			}
			t.Fatal()
		})
	}
}
