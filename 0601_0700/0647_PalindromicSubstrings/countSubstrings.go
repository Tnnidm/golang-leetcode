package countSubstrings

func countSubstrings(s string) int {
	sLen := len(s)
	result := sLen
	// 奇数子串
	for i := 1; i < sLen-1; i++ {
		for j := 1; i-j >= 0 && i+j < sLen && s[i-j] == s[i+j]; j++ {
			result++
		}
	}
	// 偶数子串
	for i := 0; i < sLen-1; i++ {
		for j := 0; i-j >= 0 && i+1+j < sLen && s[i-j] == s[i+1+j]; j++ {
			result++
		}
	}
	return result
}
