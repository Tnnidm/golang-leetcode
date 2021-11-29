package uncommonFromSentences

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	sMap := map[string]int{}
	s1List := strings.Fields(s1)
	for _, word := range s1List {
		sMap[word]++
	}
	s2List := strings.Fields(s2)
	for _, word := range s2List {
		sMap[word]++
	}
	result := []string{}
	for k, v := range sMap {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}
