package toGoatLatin

import "strings"

func toGoatLatin(sentence string) string {
	vowelMap := map[byte]bool{
		'a': true,
		'A': true,
		'e': true,
		'E': true,
		'i': true,
		'I': true,
		'o': true,
		'O': true,
		'u': true,
		'U': true,
	}
	words := strings.Fields(sentence)
	for i := range words {
		wordByte := []byte(words[i])
		if _, ok := vowelMap[wordByte[0]]; !ok {
			wordByte = append(wordByte[1:], wordByte[0])
		}
		wordByte = append(wordByte, 'm')
		for j := 0; j <= i+1; j++ {
			wordByte = append(wordByte, 'a')
		}
		words[i] = string(wordByte)
	}
	result := strings.Builder{}
	for i := range words {
		result.WriteString(words[i])
		if i != len(words)-1 {
			result.WriteRune(' ')
		}
	}
	return result.String()
}
