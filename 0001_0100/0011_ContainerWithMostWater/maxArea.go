package maxArea

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	if left == right-1 {
		return min(height[left], height[right])
	}
	var area, maxarea int
	for left < right {
		if height[left] <= height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if area > maxarea {
			maxarea = area
		}
	}
	return maxarea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
