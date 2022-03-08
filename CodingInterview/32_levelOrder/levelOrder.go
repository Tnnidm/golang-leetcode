package levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderPlain(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		thisLayer := [](*TreeNode){root}
		for len(thisLayer) != 0 {
			result = append(result, thisLayer[0].Val)
			if thisLayer[0].Left != nil {
				thisLayer = append(thisLayer, thisLayer[0].Left)
			}
			if thisLayer[0].Right != nil {
				thisLayer = append(thisLayer, thisLayer[0].Right)
			}
			thisLayer = thisLayer[1:]
		}
	}
	return result
}

func levelOrderHierarchy(root *TreeNode) [][]int {
	result := [][]int{}
	if root != nil {
		thisLayer := [](*TreeNode){root}
		for len(thisLayer) != 0 {
			thisLayerValue := []int{}
			nextLayer := [](*TreeNode){}
			for _, node := range thisLayer {
				thisLayerValue = append(thisLayerValue, node.Val)
				if node.Left != nil {
					nextLayer = append(nextLayer, node.Left)
				}
				if node.Right != nil {
					nextLayer = append(nextLayer, node.Right)
				}
			}
			result = append(result, thisLayerValue)
			thisLayer = nextLayer
		}
	}
	return result
}

func levelOrderZType(root *TreeNode) [][]int {
	result := [][]int{}
	if root != nil {
		Left2Right := true
		thisLayerNode := [](*TreeNode){root}
		for len(thisLayerNode) != 0 {
			thisLayerValue := []int{}
			nextLayerNode := [](*TreeNode){}
			if Left2Right {
				for i := len(thisLayerNode) - 1; i >= 0; i-- {
					thisLayerValue = append(thisLayerValue, thisLayerNode[i].Val)
					if thisLayerNode[i].Left != nil {
						nextLayerNode = append(nextLayerNode, thisLayerNode[i].Left)
					}
					if thisLayerNode[i].Right != nil {
						nextLayerNode = append(nextLayerNode, thisLayerNode[i].Right)
					}
				}
			} else {
				for i := len(thisLayerNode) - 1; i >= 0; i-- {
					thisLayerValue = append(thisLayerValue, thisLayerNode[i].Val)
					if thisLayerNode[i].Right != nil {
						nextLayerNode = append(nextLayerNode, thisLayerNode[i].Right)
					}
					if thisLayerNode[i].Left != nil {
						nextLayerNode = append(nextLayerNode, thisLayerNode[i].Left)
					}
				}
			}
			result = append(result, thisLayerValue)
			Left2Right = !Left2Right
			thisLayerNode = nextLayerNode
		}
	}
	return result
}
