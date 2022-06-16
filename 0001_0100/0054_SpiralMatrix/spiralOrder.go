package spiralOrder

func spiralOrder(matrix [][]int) []int {
	rowNum := len(matrix)
	colNum := len(matrix[0])
	totalNum := rowNum * colNum
	result := make([]int, totalNum)
	rowUpper, rowLower := 0, rowNum-1
	colLeft, colRight := 0, colNum-1
	for index := 0; ; {
		if colLeft < colRight {
			for i := colLeft; i <= colRight; i++ {
				result[index] = matrix[rowUpper][i]
				index++
			}
			rowUpper++
			if index == totalNum {
				return result
			}
		}
		if rowUpper < rowLower {
			for i := rowUpper; i <= rowLower; i++ {
				result[index] = matrix[i][colRight]
				index++
			}
			colRight--
			if index == totalNum {
				return result
			}
		}
		if colLeft < colRight {
			for i := colRight; i >= colLeft; i-- {
				result[index] = matrix[rowLower][i]
				index++
			}
			rowLower--
			if index == totalNum {
				return result
			}
		}
		if rowUpper <= rowLower {
			for i := rowLower; i >= rowUpper; i-- {
				result[index] = matrix[i][colLeft]
				index++
			}
			colLeft++
			if index == totalNum {
				return result
			}
		}
	}
}
