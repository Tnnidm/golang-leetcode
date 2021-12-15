package oddCells

func oddCells(m int, n int, indices [][]int) int {
	row := make([]int, m)
	col := make([]int, n)
	for i := 0; i < len(indices); i++ {
		row[indices[i][0]]++
		col[indices[i][1]]++
	}
	rowOddNum := 0
	for i := 0; i < m; i++ {
		rowOddNum = rowOddNum + row[i]%2
	}
	colOddNum := 0
	for i := 0; i < n; i++ {
		colOddNum = colOddNum + col[i]%2
	}
	return rowOddNum*(n-colOddNum) + colOddNum*(m-rowOddNum)
}
