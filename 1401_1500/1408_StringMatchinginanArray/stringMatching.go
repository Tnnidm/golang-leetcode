package stringMatching

import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	wordLen := len(words)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) <= len(words[j])
	})
	result := []string{}
	for i := 0; i < wordLen; i++ {
		for j := i + 1; j < wordLen; j++ {
			if i != j && strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}
	return result
}
