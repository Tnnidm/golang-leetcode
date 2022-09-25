package minDistance

// // 递归做法，但是超时了
// func minDistance(word1 string, word2 string) int {
// 	word1Len := len(word1)
// 	word2Len := len(word2)
// 	if word1Len < word2Len {
// 		return minDistance(word2, word1)
// 	}
// 	if word1Len == 0 || word2Len == 0 {
// 		return word1Len + word2Len
// 	}
// 	result := 1 + minDistance(word1, word2[1:])               // 在word1头部加上word2[0]
// 	result = min(result, 1+minDistance(word1[1:], word2[1:])) // 把word1[0]换成word2[0]
// 	for i := range word1 {
// 		if word1[i] == word2[0] {
// 			result = min(result, i+minDistance(word1[i+1:], word2[1:])) // 把word1[:i]删除
// 		}
// 	}
// 	return result
// }

// 动态规划做法
func minDistance(word1 string, word2 string) int {
	word1Len, word2Len := len(word1), len(word2)
	dp := make([]int, word1Len+1)
	for i := range dp {
		dp[i] = word1Len - i
	}
	for j := word2Len - 1; j >= 0; j-- {
		for i := 0; i < word1Len; i++ {
			dp[i] = min(dp[i]+1, dp[i+1]+1)
			for k := i; k < word1Len; k++ {
				if word1[k] == word2[j] {
					dp[i] = min(dp[i], k-i+dp[k+1])
				}
			}
		}
		dp[word1Len] = word2Len - j
	}
	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
