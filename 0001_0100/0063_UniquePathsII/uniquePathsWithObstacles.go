package uniquePathsWithObstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	dp := make([]int, col)
	for j := 0; j < col && obstacleGrid[0][j] == 0; j++ {
		dp[j] = 1
	}
	for i := 1; i < row; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[col-1]
}
