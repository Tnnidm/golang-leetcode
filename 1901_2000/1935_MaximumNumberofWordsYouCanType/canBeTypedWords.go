package canBeTypedWords

import "strings"

func canBeTypedWords(text, brokenLetters string) int {
	result := 0
	for _, word := range strings.Split(text, " ") {
		if !strings.ContainsAny(word, brokenLetters) {
			result++
		}
	}
	return result
}
