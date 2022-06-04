package maximalSquare

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n)
	var prev, result int
	for i := 0; i < m; i++ {
		prev = dp[0]
		dp[0] = int(matrix[i][0] - '0')
		if dp[0] > result {
			result = dp[0]
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[j], prev = min(min(dp[j-1], dp[j]), prev)+1, dp[j]
				if dp[j] > result {
					result = dp[j]
				}
			} else {
				dp[j] = 0
			}
		}
	}
	return result * result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
