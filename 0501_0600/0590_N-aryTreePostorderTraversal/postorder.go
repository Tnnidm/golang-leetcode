package postorder

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	result := []int{}
	if root != nil {
		stack := []*Node{root}
		stackLen := len(stack)
		var node *Node
		for stackLen != 0 {
			node = stack[stackLen-1]
			stack = stack[:stackLen-1]
			result = append(result, node.Val)
			for _, v := range node.Children {
				stack = append(stack, v)
			}
			stackLen = len(stack)
		}
		resultLen := len(result)
		for i := 0; i < resultLen/2; i++ {
			result[i], result[resultLen-1-i] = result[resultLen-1-i], result[i]
		}
	}
	return result
}
