package largestRectangleArea

func largestRectangleArea(heights []int) int {
	heightsNum := len(heights)
	stack := []int{-1}
	stackLen := 1
	maxArea := 0
	for index, height := range heights {
		for stack[stackLen-1] != -1 && heights[stack[stackLen-1]] >= height {
			stackLen--
			topIndex := stack[stackLen]
			stack = stack[:stackLen]
			maxArea = max(maxArea, heights[topIndex]*(index-stack[stackLen-1]-1))
		}
		stack = append(stack, index)
		stackLen++
	}
	for stack[stackLen-1] != -1 {
		stackLen--
		topIndex := stack[stackLen]
		stack = stack[:stackLen]
		maxArea = max(maxArea, heights[topIndex]*(heightsNum-stack[stackLen-1]-1))
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
