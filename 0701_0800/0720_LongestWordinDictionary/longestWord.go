package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	wordMap := make(map[string]struct{})
	longestLength := 0
	longestWord := ""
	for _, word := range words {
		wordLen := len(word)
		if _, ok := wordMap[word[:wordLen-1]]; wordLen == 1 || ok {
			wordMap[word] = struct{}{}
			if wordLen > longestLength {
				longestLength = wordLen
				longestWord = word
			}
		}
	}
	return longestWord
}

func main() {
	words := []string{"rac", "rs", "ra", "on", "r", "otif", "o", "onpdu", "rsf", "rs", "ot", "oti", "racy", "onpd"}
	fmt.Println(longestWord(words))
}
