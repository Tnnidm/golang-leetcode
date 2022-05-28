package rotate

func rotate(matrix [][]int) {
	n := len(matrix)
	var length, width int
	if n%2 == 1 {
		length = n>>1 + 1
		width = n >> 1
	} else {
		length = n >> 1
		width = n >> 1
	}
	var temp int
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			temp = matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}
