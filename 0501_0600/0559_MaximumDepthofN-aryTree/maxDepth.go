package maxDepth

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	} else {
		maxChildDepth := 0
		childDepth := 0
		for i := 0; i < len(root.Children); i++ {
			childDepth = maxDepth(root.Children[i])
			if childDepth > maxChildDepth {
				maxChildDepth = childDepth
			}
		}
		return maxChildDepth + 1
	}
}
