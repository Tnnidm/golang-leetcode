package mergeAlternately

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	word1Len := len(word1)
	word2Len := len(word2)
	result := strings.Builder{}
	if word1Len <= word2Len {
		for i := range word1 {
			result.WriteByte(word1[i])
			result.WriteByte(word2[i])
		}
		for i := word1Len; i < word2Len; i++ {
			result.WriteByte(word2[i])
		}
	} else {
		for i := range word2 {
			result.WriteByte(word1[i])
			result.WriteByte(word2[i])
		}
		for i := word2Len; i < word1Len; i++ {
			result.WriteByte(word1[i])
		}
	}
	return result.String()
}
