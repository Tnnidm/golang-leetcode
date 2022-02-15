package buildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	nodeNum := len(preorder)
	if nodeNum == 0 {
		return nil
	}
	headVal := preorder[0]
	headIndex := 0
	for i, v := range inorder {
		if v == headVal {
			headIndex = i
		}
	}
	leftNum := headIndex
	rightNum := nodeNum - 1 - headIndex
	return &TreeNode{headVal, buildTree(preorder[1:1+leftNum], inorder[:leftNum]), buildTree(preorder[nodeNum-rightNum:], inorder[nodeNum-rightNum:])}
}
