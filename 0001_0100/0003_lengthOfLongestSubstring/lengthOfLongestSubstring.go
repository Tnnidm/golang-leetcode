package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	lastIndex := [128]int{}
	var left, maxLength int
	for i := 0; i < len(s); i++ {
		left = max(left, lastIndex[int(s[i])])
		maxLength = max(maxLength, i-left+1)
		lastIndex[int(s[i])] = i + 1
	}
	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
