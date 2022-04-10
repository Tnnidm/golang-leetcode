package isPalindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isValid(s[left]) {
			left++
		}
		for left < right && !isValid(s[right]) {
			right--
		}
		// 这里使用局部的unicode.ToLower()会比一开始在全局strings.ToLower()消耗的时间空间都少很多
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isValid(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
