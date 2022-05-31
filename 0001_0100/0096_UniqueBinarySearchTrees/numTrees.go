package numTrees

func numTrees(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		for c := 1; c <= i; c++ {
			dp[i] += dp[c-1] * dp[i-c]
		}
	}
	return dp[n]
}
