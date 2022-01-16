package countGoodSubstrings

func countGoodSubstrings(s string) int {
	sLen := len(s)
	if sLen < 3 {
		return 0
	}
	result := 0
	for i := 0; i <= sLen-3; i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			result++
		}
	}
	return result
}
