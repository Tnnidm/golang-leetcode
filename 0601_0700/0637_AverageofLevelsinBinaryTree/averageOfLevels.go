package averageOfLevels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	thisLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}
	sum := 0
	result := []float64{}
	for len(thisLevel) != 0 {
		for _, node := range thisLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		result = append(result, float64(sum)/float64((len(thisLevel))))
		sum = 0
		thisLevel = nextLevel
		nextLevel = []*TreeNode{}
	}
	return result
}
