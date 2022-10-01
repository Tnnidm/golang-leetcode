package longestContinuousSubstring

func longestContinuousSubstring(s string) int {
	length := 1
	maxLen := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			length++
			if length > maxLen {
				maxLen = length
			}
		} else {
			length = 1
		}
	}
	return maxLen
}
