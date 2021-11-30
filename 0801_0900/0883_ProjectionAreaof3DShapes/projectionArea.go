package projectionArea

func projectionArea(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	result := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != 0 {
				result++
			}
		}
	}
	var maxHigh int
	for i := 0; i < row; i++ {
		maxHigh = 0
		for j := 0; j < col; j++ {
			if grid[i][j] > maxHigh {
				maxHigh = grid[i][j]
			}
		}
		result += maxHigh
	}
	for j := 0; j < col; j++ {
		maxHigh = 0
		for i := 0; i < row; i++ {
			if grid[i][j] > maxHigh {
				maxHigh = grid[i][j]
			}
		}
		result += maxHigh
	}
	return result
}
