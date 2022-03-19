package reverseWords

import (
	"strings"
)

// // 方法1
// func reverseWords(s string) string {
// 	words := strings.Fields(s)
// 	wordNum := len(words)
// 	result := strings.Builder{}
// 	for i := wordNum - 1; i >= 0; i-- {
// 		result.WriteString(words[i])
// 		if i != 0 {
// 			result.WriteString(" ")
// 		}
// 	}
// 	return result.String()
// }

// 方法2
func reverseWords(s string) string {
	words := strings.Fields(s)
	wordNum := len(words)
	for i := 0; i <= wordNum/2; i++ {
		words[i], words[wordNum-1-i] = words[wordNum-1-i], words[i]
	}
	return strings.Join(words, " ")
}
