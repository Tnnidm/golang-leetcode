package firstPalindrome

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindromic(word) {
			return word
		}
	}
	return ""
}

func isPalindromic(word string) bool {
	wordLen := len(word)
	for i := 0; i < wordLen/2; i++ {
		if word[i] != word[wordLen-1-i] {
			return false
		}
	}
	return true
}
