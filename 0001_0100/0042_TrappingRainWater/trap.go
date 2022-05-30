package trap

// func trap(height []int) int {
// 	heightLen := len(height)
// 	rightMax := make([]int, heightLen)
// 	rightMax[heightLen-1] = height[heightLen-1]
// 	for i := heightLen - 2; i >= 0; i-- {
// 		rightMax[i] = max(rightMax[i+1], height[i])
// 	}
// 	leftMax := height[0]
// 	result := 0
// 	for i := 1; i < heightLen-1; i++ {
// 		leftMax = max(leftMax, height[i])
// 		result += min(leftMax, rightMax[i]) - height[i]
// 	}
// 	return result
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// 动态规划法
func trap(arr []int) int {
	left, right := 0, len(arr)-1
	leftMax, rightMax := arr[left], arr[right]
	result := 0
	for left < right { // left+1位置瓶颈来自左边，left+1位置蓄水量可以求出来
		if leftMax <= rightMax {
			if arr[left+1] < leftMax {
				result += leftMax - arr[left+1]
			} else {
				leftMax = arr[left+1]
			}
			left++
		} else { // right-1位置瓶颈来自右边，right-1位置蓄水量可以求出来
			if arr[right-1] < rightMax {
				result += rightMax - arr[right-1]
			} else {
				rightMax = arr[right-1]
			}
			right--
		}
	}
	return result
}
