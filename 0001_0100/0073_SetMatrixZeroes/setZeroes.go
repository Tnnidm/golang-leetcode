package setZeroes

func setZeroes(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	firstRowIsZeros, firstColIsZeros := false, false
	for j := 0; j < col; j++ {
		if matrix[0][j] == 0 {
			firstRowIsZeros = true
			break
		}
	}
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			firstColIsZeros = true
			break
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < row; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < col; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < col; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < row; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if firstRowIsZeros {
		for j := 0; j < col; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColIsZeros {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}
