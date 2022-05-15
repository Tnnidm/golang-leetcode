package minPathSum

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([]int, col)
	dp[0] = grid[0][0]
	for j := 1; j < col; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < row; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < col; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[col-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
