package preorder

type Node struct {
	Val      int
	Children []*Node
}

// 递归解法
// func preorder(root *Node) []int {
// 	result := []int{}
// 	N_preorder(root, &result)
// 	return result
// }

// func N_preorder(root *Node, node *[]int) {
// 	if root != nil {
// 		*node = append(*node, root.Val)
// 		for _, v := range root.Children {
// 			N_preorder(v, node)
// 		}
// 	}
// }

// 迭代解法
func preorder(root *Node) []int {
	result := []int{}
	if root != nil {
		stack := []*Node{root}
		stackLen := len(stack)
		var node *Node
		for stackLen != 0 {
			node = stack[stackLen-1]
			stack = stack[:stackLen-1]
			result = append(result, node.Val)
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
			stackLen = len(stack)
		}
	}
	return result
}
