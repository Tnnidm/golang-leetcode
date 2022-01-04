package countVowelSubstrings

import "strings"

func countVowelSubstrings(word string) int {
	wordLen := len(word)
	if wordLen < 5 {
		return 0
	}
	result := 0
	for i := 0; i <= wordLen-5; i++ {
		for j := i + 5; j <= wordLen; j++ {
			if isVowelSubstring(word[i:j]) {
				result++
			}
		}
	}
	return result
}

func isVowelSubstring(subWord string) bool {
	for _, v := range subWord {
		if v != 'a' && v != 'e' && v != 'i' && v != 'o' && v != 'u' {
			return false
		}
	}
	if strings.ContainsRune(subWord, 'a') && strings.ContainsRune(subWord, 'e') && strings.ContainsRune(subWord, 'i') && strings.ContainsRune(subWord, 'o') && strings.ContainsRune(subWord, 'u') {
		return true
	}
	return false
}
