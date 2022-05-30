package wordBreak

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}
	sLen := len(s)
	dp := make([]bool, sLen+1)
	dp[0] = true
	for i := 1; i <= sLen; i++ {
		for j := 0; j < i && !dp[i]; j++ {
			dp[i] = dp[j] && wordMap[s[j:i]]
		}
	}
	return dp[sLen]
}
