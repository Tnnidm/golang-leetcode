package sortSentence

import (
	"sort"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Fields(s)
	sort.Slice(words, func(i, j int) bool {
		return words[i][len(words[i])-1] <= words[j][len(words[j])-1]
	})
	result := strings.Builder{}
	for i, v := range words {
		result.WriteString(v[:len(v)-1])
		if i != len(words)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}
