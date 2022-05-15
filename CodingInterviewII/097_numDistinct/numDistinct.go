package numDistinct

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	dp := make([]int, tLen+1)
	for i := 1; i <= sLen; i++ {
		dp[0] = 1
		for j := min(i, tLen); j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[tLen]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
