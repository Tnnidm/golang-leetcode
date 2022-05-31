package buildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	headNum := preorder[0]
	headIndex := 0
	for index, num := range inorder {
		if num == headNum {
			headIndex = index
			break
		}
	}
	return &TreeNode{headNum, buildTree(preorder[1:1+headIndex], inorder[0:headIndex]), buildTree(preorder[1+headIndex:], inorder[headIndex+1:])}
}
