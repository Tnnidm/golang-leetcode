package minimumTotal

func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	dp := make([]int, row)
	dp[0] = triangle[0][0]
	for i := 1; i < row; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] = dp[0] + triangle[i][0]
	}
	result := 10000
	for j := 0; j < row; j++ {
		if dp[j] < result {
			result = dp[j]
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
