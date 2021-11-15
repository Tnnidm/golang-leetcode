package imageSmoother

func imageSmoother(img [][]int) [][]int {
	xLen := len(img)
	yLen := len(img[0])
	var sum, count int

	result := make([][]int, xLen)
	for x := range result {
		result[x] = make([]int, yLen)
	}

	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			sum, count = 0, 0
			for i := max(0, x-1); i <= min(x+1, xLen-1); i++ {
				for j := max(0, y-1); j <= min(y+1, yLen-1); j++ {
					sum += img[i][j]
					count++
				}
			}
			result[x][y] = sum / count
		}
	}
	return result
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
