package mostCommonWord

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	// paragraphSplit := strings.Split(strings.ToLower(paragraph), " ")
	// 使用非字母来分割原句子，strings包值得好好看一下
	paragraphSplit := strings.FieldsFunc(strings.ToLower(paragraph), func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	paraMap := make(map[string]int)
	maxWord := ""
	for _, v := range paragraphSplit {
		paraMap[v]++
		if paraMap[v] > paraMap[maxWord] {
			if !isBanned(v, banned) {
				maxWord = v
			}
		}
	}
	return maxWord
}

func isBanned(word string, banned []string) bool {
	for _, v := range banned {
		if word == v {
			return true
		}
	}
	return false
}
