package countNegatives

func countNegatives(grid [][]int) int {
	rowNum := len(grid)
	colNum := len(grid[0])
	col := colNum
	result := 0
	for i := 0; i < rowNum; i++ {
		for col != 0 && grid[i][col-1] < 0 {
			col--
		}
		result = result + colNum - col
	}
	return result
}
