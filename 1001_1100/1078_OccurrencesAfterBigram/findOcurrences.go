package findOcurrences

import "strings"

func findOcurrences(text string, first string, second string) []string {
	words := strings.Fields(text)
	wordNum := len(words)
	if wordNum < 3 {
		return []string{}
	}
	result := []string{}
	for i := 0; i < wordNum-2; i++ {
		if words[i] == first && words[i+1] == second {
			result = append(result, words[i+2])
		}
	}
	return result
}
