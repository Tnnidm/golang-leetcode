package postorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归方法
// func postorderTraversal(root *TreeNode) []int {
// 	result := []int{}
// 	postorder(root, &result)
// 	return result
// }

// func postorder(root *TreeNode, result *[]int) {
// 	if root != nil {
// 		postorder(root.Left, result)
// 		postorder(root.Right, result)
// 		*result = append(*result, root.Val)
// 	}
// }

// 迭代方法
type Stack struct {
	volume [](*TreeNode)
	length int
}

func (s *Stack) Push(value *TreeNode) {
	s.volume = append(s.volume, value)
	s.length++
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) Top() *TreeNode {
	return s.volume[s.length-1]
}

func (s *Stack) Pop() {
	s.length--
	s.volume = s.volume[:s.length]
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := Stack{}
	var cur, prev *TreeNode
	cur = root
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		cur = stack.Top()
		if cur.Right != nil && cur.Right != prev {
			cur = cur.Right
		} else {
			stack.Pop()
			result = append(result, cur.Val)
			prev = cur
			cur = nil
		}
	}
	return result
}
