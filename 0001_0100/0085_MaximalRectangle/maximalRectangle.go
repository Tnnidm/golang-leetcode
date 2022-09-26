package maximalRectangle

func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}
	hatch := make([]int, cols)
	result := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				hatch[j]++
			} else {
				hatch[j] = 0
			}
		}
		result = max(result, maximalRectangleInHatch(hatch))
	}
	return result
}

func maximalRectangleInHatch(hatch []int) int {
	stack := []int{-1}
	stackLen := 1
	result := hatch[0]
	for i := range hatch {
		for stackLen > 1 && hatch[i] <= hatch[stack[stackLen-1]] {
			result = max(result, hatch[stack[stackLen-1]]*(i-stack[stackLen-2]-1))
			stack = stack[:stackLen-1]
			stackLen--
		}
		stack = append(stack, i)
		stackLen++
	}
	for i := len(stack) - 1; i >= 1; i-- {
		result = max(result, hatch[stack[i]]*(len(hatch)-stack[i-1]-1))
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
