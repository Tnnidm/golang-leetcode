package shiftGrid

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	queue := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			queue[i*n+j] = grid[i][j]
		}
	}
	k = k % (m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = queue[(i*n+j+m*n-k)%(m*n)]
		}
	}
	return grid
}
