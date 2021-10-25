package lowestCommonAncestor

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tNode := &TreeNode{}
	tNode.Val = 6

	tNode1 := &TreeNode{}
	tNode1.Val = 2

	tNode2 := &TreeNode{}
	tNode2.Val = 8

	tNode3 := &TreeNode{}
	tNode3.Val = 0

	tNode4 := &TreeNode{}
	tNode4.Val = 4

	tNode5 := &TreeNode{}
	tNode5.Val = 7

	tNode6 := &TreeNode{}
	tNode6.Val = 9

	tNode9 := &TreeNode{}
	tNode9.Val = 3

	tNode10 := &TreeNode{}
	tNode10.Val = 5

	// 第一层
	tNode.Left = tNode1
	tNode.Right = tNode2

	// 第二层
	tNode1.Left = tNode3
	tNode1.Right = tNode4
	tNode2.Left = tNode5
	tNode2.Right = tNode6

	// 第三层
	tNode4.Left = tNode9
	tNode4.Right = tNode10

	luckyNode := lowestCommonAncestor(tNode, tNode1, tNode2)
	t.Logf("luckyNode.value = %d", luckyNode.Val)
	if luckyNode != tNode {
		t.Fatalf("Unlucky, false.")
	}
}
