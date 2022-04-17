package CBTInserter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root  *TreeNode
	queue [](*TreeNode)
}

func Constructor(root *TreeNode) CBTInserter {
	ci := CBTInserter{root, [](*TreeNode){root}}
	for isComplete(ci.queue[0]) {
		ci.queue = append(ci.queue, ci.queue[0].Left)
		ci.queue = append(ci.queue, ci.queue[0].Right)
		ci.queue = ci.queue[1:]
	}
	return ci
}

func isComplete(root *TreeNode) bool {
	return root.Left != nil && root.Right != nil
}

func (ci *CBTInserter) Insert(v int) int {
	temp := ci.queue[0].Val
	if ci.queue[0].Left == nil {
		ci.queue[0].Left = &TreeNode{v, nil, nil}
	} else {
		ci.queue[0].Right = &TreeNode{v, nil, nil}
		ci.queue = append(ci.queue, ci.queue[0].Left)
		ci.queue = append(ci.queue, ci.queue[0].Right)
		ci.queue = ci.queue[1:]
	}
	return temp
}

func (ci *CBTInserter) Get_root() *TreeNode {
	return ci.root
}
