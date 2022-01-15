package numOfStrings

import "strings"

func numOfStrings(patterns []string, word string) int {
	result := 0
	for i := range patterns {
		if strings.Contains(word, patterns[i]) {
			result++
		}
	}
	return result
}
