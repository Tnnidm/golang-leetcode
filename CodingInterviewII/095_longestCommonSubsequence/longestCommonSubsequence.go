package longestCommonSubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	if len1 < len2 {
		return longestCommonSubsequence(text2, text1)
	}
	dp := make([]int, len2+1)
	var prev, cur int
	for i := 0; i < len1; i++ {
		prev = dp[0]
		for j := 0; j < len2; j++ {
			if text1[i] == text2[j] {
				cur = prev + 1
			} else {
				cur = max(dp[j+1], dp[j])
			}
			prev = dp[j+1]
			dp[j+1] = cur
		}
	}
	return dp[len2]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
