package transpose

func transpose(matrix [][]int) [][]int {
	row := len(matrix)
	col := len(matrix[0])
	result := make([][]int, col)
	for i := 0; i < col; i++ {
		result[i] = make([]int, row)
	}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}
