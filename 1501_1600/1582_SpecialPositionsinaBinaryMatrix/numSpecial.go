package numSpecial

func numSpecial(mat [][]int) int {
	result := 0
	rows := len(mat)
	cols := len(mat[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 {
				sum := 0
				for m := 0; m < cols; m++ {
					sum = sum + mat[i][m]
				}
				for n := 0; n < rows; n++ {
					sum = sum + mat[n][j]
				}
				if sum == 2 {
					result++
				}
			}
		}
	}
	return result
}
