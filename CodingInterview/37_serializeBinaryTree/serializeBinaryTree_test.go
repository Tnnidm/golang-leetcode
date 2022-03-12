package serializeBinaryTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	cases := []struct {
		name    string
		input   *TreeNode
		expect1 []string
		expect2 *TreeNode
	}{
		{
			"case 1",
			&TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}},
			[]string{"4", "2", "1", "null", "null", "3", "null", "null", "5", "null", "null"},
			&TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret1 := serializeBinaryTree(c.input)
			if !reflect.DeepEqual(ret1, c.expect1) {
				t.Fatalf("expect: %v, but got %v, with input: %v", c.expect1, ret1, c.input)
			}
			ret2 := deserializeBinaryTree(ret1)
			fmt.Println(serializeBinaryTree(ret2))
			if !reflect.DeepEqual(ret2, c.expect2) {
				t.Fatalf("expect: %v, but got %v, with input: %v", c.expect2, ret2, c.expect1)
			}
		})
	}
}
