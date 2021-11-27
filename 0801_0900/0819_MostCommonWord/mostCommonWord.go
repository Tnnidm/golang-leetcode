package mostCommonWord

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	// 使用非字母来分割原句子，strings包值得好好看一下
	words := strings.FieldsFunc(paragraph, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	mapWords := make(map[string]int)

	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		mapWords[word]++
	}

	max := 0
	res := ""
	for word, count := range mapWords {
		if count > max && !inBanned(word, banned) {
			max = count
			res = word
		}
	}

	return res
}

func inBanned(word string, banned []string) bool {
	for _, bannedStr := range banned {
		if word == bannedStr {
			return true
		}
	}
	return false
}
