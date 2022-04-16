package largestRectangleArea

func largestRectangleArea(heights []int) int {
	heightsLen := len(heights)
	stack := make([]int, 0, 16)
	stackLen := 0
	maxArea := 0
	for i := 0; i < heightsLen; i++ {
		for stackLen > 0 && heights[stack[stackLen-1]] >= heights[i] {
			if stackLen == 1 {
				maxArea = max(maxArea, heights[stack[stackLen-1]]*i)
			} else {
				maxArea = max(maxArea, heights[stack[stackLen-1]]*(i-1-stack[stackLen-2]))
			}
			stack = stack[:stackLen-1]
			stackLen--
		}
		stack = append(stack, i)
		stackLen++
	}
	for i := stackLen - 1; i >= 1; i-- {
		maxArea = max(maxArea, heights[stack[i]]*(heightsLen-1-stack[i-1]))
	}
	maxArea = max(maxArea, heights[stack[0]]*heightsLen)
	return maxArea
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
