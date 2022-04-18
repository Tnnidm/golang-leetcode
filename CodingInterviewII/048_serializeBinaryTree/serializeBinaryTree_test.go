package serializeBinaryTree

import (
	"reflect"
	"testing"
)

func Test_serializeBinaryTree(t *testing.T) {
	cases := []struct {
		name  string
		input *TreeNode
	}{
		{"case 1", &TreeNode{1, &TreeNode{0, &TreeNode{0, nil, nil}, nil}, &TreeNode{1, nil, &TreeNode{0, nil, nil}}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			obj := Constructor()
			ret1 := obj.serialize(c.input)
			ret2 := obj.deserialize(ret1)
			if !reflect.DeepEqual(ret2, c.input) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.input, ret2, c.input)
			}
		})
	}
}
