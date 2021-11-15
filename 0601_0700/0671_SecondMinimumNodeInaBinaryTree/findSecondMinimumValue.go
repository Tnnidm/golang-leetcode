package findSecondMinimumValue

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 解法1
// func findSecondMinimumValue(root *TreeNode) int {
// 	arr := []int{}
// 	findArr(root, root.Val, &arr)
// 	if len(arr) > 0 {
// 		minValue := arr[0]
// 		for _, v := range arr {
// 			if v < minValue {
// 				minValue = v
// 			}
// 		}
// 		return minValue
// 	} else {
// 		return -1
// 	}
// }

// func findArr(root *TreeNode, minValue int, arr *[]int) {
// 	if root.Val == minValue {
// 		if root.Left != nil {
// 			findArr(root.Left, minValue, arr)
// 		}
// 		if root.Right != nil {
// 			findArr(root.Right, minValue, arr)
// 		}
// 	} else {
// 		*arr = append(*arr, root.Val)
// 	}
// }

// 解法2
func findSecondMinimumValue(root *TreeNode) int {
	secondMin := math.MaxInt64
	findArr(root, root.Val, &secondMin)
	if secondMin == math.MaxInt64 {
		return -1
	} else {
		return secondMin
	}
}

func findArr(root *TreeNode, minValue int, secondMin *int) {
	if root.Val == minValue {
		if root.Left != nil {
			findArr(root.Left, minValue, secondMin)
		}
		if root.Right != nil {
			findArr(root.Right, minValue, secondMin)
		}
	} else {
		if root.Val < *secondMin {
			*secondMin = root.Val
		}
	}
}
