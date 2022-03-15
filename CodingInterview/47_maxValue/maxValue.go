package maxValue

func maxValue(grid [][]int) int {
	gridRow := len(grid)
	gridCol := len(grid[0])
	for j := 1; j < gridCol; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < gridRow; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < gridRow; i++ {
		for j := 1; j < gridCol; j++ {
			grid[i][j] += max(grid[i][j-1], grid[i-1][j])
		}
	}
	return grid[gridRow-1][gridCol-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
