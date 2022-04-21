package BSTIterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack [](*TreeNode)
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{[](*TreeNode){}, root}
	for iter.cur != nil {
		iter.stack = append(iter.stack, iter.cur)
		iter.cur = iter.cur.Left
	}
	return iter
}

func (iter *BSTIterator) Next() int {
	if iter.cur == nil && len(iter.stack) > 0 {
		stackTop := len(iter.stack) - 1
		iter.cur = iter.stack[stackTop]
		iter.stack = iter.stack[:stackTop]
	}
	result := iter.cur.Val
	iter.cur = iter.cur.Right
	for iter.cur != nil {
		iter.stack = append(iter.stack, iter.cur)
		iter.cur = iter.cur.Left
	}
	return result
}

func (iter *BSTIterator) HasNext() bool {
	return iter.cur != nil || len(iter.stack) != 0
}
