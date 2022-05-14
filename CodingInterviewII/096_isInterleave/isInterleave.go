package isInterleave

func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len < s2Len {
		return isInterleave(s2, s1, s3)
	}
	s3Len := len(s3)
	// 特例处理
	if s1Len+s2Len != s3Len {
		return false
	}
	if s1Len == 0 {
		return s2 == s3
	}
	if s2Len == 0 {
		return s1 == s3
	}
	// 检查交织
	dp := make([]bool, s2Len+1) // dp[i][j]表示s1[:i]和s2[:j]可以交织成为s3[:i+j]
	dp[0] = true
	// 没有s1的情况
	for j := 0; j < s2Len; j++ {
		dp[j+1] = dp[j] && s2[j] == s3[j]
	}
	// 有s1的情况
	for i := 1; i <= s1Len; i++ {
		// 表示s1[:i]参与了交织
		dp[0] = s1[:i] == s3[:i]
		for j := 1; j <= s2Len; j++ {
			// 表示s2[:j]参与了交织
			dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1] || dp[j] && s1[i-1] == s3[i+j-1]
		}
	}
	return dp[s2Len]
}
