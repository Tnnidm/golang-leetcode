package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	lastIndex := make([]int, 256)
	left := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		left = max(left, lastIndex[s[i]])
		maxLen = max(maxLen, i-left+1)
		lastIndex[s[i]] = i + 1
	}
	return maxLen
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
