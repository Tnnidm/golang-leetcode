package luckyNumbers

func luckyNumbers(matrix [][]int) []int {
	rowNum := len(matrix)
	colNUm := len(matrix[0])
	result := []int{}
	for i := 0; i < rowNum; i++ {
		minCol := 0
		for j := 0; j < colNUm; j++ {
			if matrix[i][j] < matrix[i][minCol] {
				minCol = j
			}
		}
		flag := true
		for k := 0; k < rowNum; k++ {
			if matrix[k][minCol] > matrix[i][minCol] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, matrix[i][minCol])
		}
	}
	return result
}
