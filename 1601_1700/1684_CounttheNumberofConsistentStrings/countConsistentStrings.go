package countConsistentStrings

import "strings"

func countConsistentStrings(allowed string, words []string) int {
	result := 0
	for _, word := range words {
		result++
		for _, char := range word {
			if !strings.ContainsRune(allowed, char) {
				result--
				break
			}
		}
	}
	return result
}
