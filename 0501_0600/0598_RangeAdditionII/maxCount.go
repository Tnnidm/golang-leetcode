package maxCount

func maxCount(m int, n int, ops [][]int) int {
	xMin, yMin := m, n
	for i := 0; i < len(ops); i++ {
		xMin = min(xMin, ops[i][0])
		yMin = min(yMin, ops[i][1])
	}
	return xMin * yMin
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
