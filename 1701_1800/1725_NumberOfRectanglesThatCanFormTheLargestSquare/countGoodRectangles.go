package countGoodRectangles

func countGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	maxLenNum := 0
	var length int
	for i := 0; i < len(rectangles); i++ {
		length = min(rectangles[i][0], rectangles[i][1])
		if length > maxLen {
			maxLen = length
			maxLenNum = 1
		} else if length == maxLen {
			maxLenNum++
		}
	}
	return maxLenNum
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
