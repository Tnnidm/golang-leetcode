package longestPalindrome

// // 方法1: 暴力搜索
// func longestPalindrome(s string) string {
// 	sLen := len(s)
// 	maxLen := 1
// 	result := s[:1]
// 	for left := 0; left < sLen-1; left++ {
// 		for right := sLen - 1; right > left; right-- {
// 			if s[left] == s[right] && isPalindrome(&s, left, right) && right-left+1 > maxLen {
// 				maxLen = right - left + 1
// 				result = s[left : right+1]
// 			}
// 		}
// 	}
// 	return result
// }

// func isPalindrome(s *string, left, right int) bool {
// 	for left < right {
// 		if (*s)[left] != (*s)[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

// 方法2: 中心扩散法
func longestPalindrome(s string) string {
	sLen := len(s)
	maxLen := 1
	result := s[:1]
	// 先找奇数回文子串
	for i := 1; i < sLen-1; i++ {
		left := i
		right := i
		for left > 0 && right < sLen-1 && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
			result = s[left : right+1]
		}
	}
	// 再找偶数回文串
	for i := 0; i < sLen-1; i++ {
		left := i
		right := i + 1
		if s[left] == s[right] {
			for left > 0 && right < sLen-1 && s[left-1] == s[right+1] {
				left--
				right++
			}
			if right-left+1 > maxLen {
				maxLen = right - left + 1
				result = s[left : right+1]
			}
		}
	}
	return result
}
