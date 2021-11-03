package islandPerimeter

func islandPerimeter(grid [][]int) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if (j == 0 || grid[i][j-1] == 0) && grid[i][j] == 1 {
				sum += 2
			}
		}
	}
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			if (i == 0 || grid[i-1][j] == 0) && grid[i][j] == 1 {
				sum += 2
			}
		}
	}
	return sum
}
