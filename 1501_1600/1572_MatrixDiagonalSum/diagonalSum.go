package diagonalSum

func diagonalSum(mat [][]int) int {
	result := 0
	rowNum := len(mat)
	for i := 0; i < rowNum; i++ {
		if i != rowNum-1-i {
			result = result + mat[i][i] + mat[i][rowNum-1-i]
		} else {
			result = result + mat[i][i]
		}
	}
	return result
}
