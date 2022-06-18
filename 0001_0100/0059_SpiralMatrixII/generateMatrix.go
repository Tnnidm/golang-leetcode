package generateMatrix

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	maxNum := n * n
	index := 1
	rowUpper, rowLower := 0, n-1
	colLeft, colRight := 0, n-1
	for index <= maxNum {
		if colLeft <= colRight {
			for j := colLeft; j <= colRight; j++ {
				result[rowUpper][j] = index
				index++
			}
			rowUpper++
			if index > maxNum {
				return result
			}
		}
		if rowUpper < rowLower {
			for i := rowUpper; i <= rowLower; i++ {
				result[i][colRight] = index
				index++
			}
			colRight--
			if index > maxNum {
				return result
			}
		}
		if colLeft < colRight {
			for j := colRight; j >= colLeft; j-- {
				result[rowLower][j] = index
				index++
			}
			rowLower--
			if index > maxNum {
				return result
			}
		}
		if rowUpper < rowLower {
			for i := rowLower; i >= rowUpper; i-- {
				result[i][colLeft] = index
				index++
			}
			colLeft++
			if index > maxNum {
				return result
			}
		}
	}
	return result
}
