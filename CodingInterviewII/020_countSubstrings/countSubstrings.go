package countSubstrings

func countSubstrings(s string) int {
	sLen := len(s)
	result := 0
	for i := 0; i < len(s); i++ {
		countPalindrome(&s, &result, &sLen, i, i)
		countPalindrome(&s, &result, &sLen, i, i+1)
	}
	return result
}

func countPalindrome(s *string, result, sLen *int, left, right int) {
	for i := 0; ; i++ {
		if left-i >= 0 && right+i < *sLen && (*s)[left-i] == (*s)[right+i] {
			*result++
		} else {
			break
		}
	}
}
