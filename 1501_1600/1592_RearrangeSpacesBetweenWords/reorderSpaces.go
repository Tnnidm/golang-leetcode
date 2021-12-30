package reorderSpaces

import "strings"

func reorderSpaces(text string) string {
	spaceNum := 0
	for _, v := range text {
		if v == ' ' {
			spaceNum++
		}
	}
	words := strings.Fields(text)
	wordNum := len(words)
	result := strings.Builder{}
	if wordNum == 1 {
		result.WriteString(words[0])
		for i := 0; i < spaceNum; i++ {
			result.WriteByte(' ')
		}
	} else {
		eachSpaceNum := spaceNum / (wordNum - 1)
		for i, v := range words {
			result.WriteString(v)
			if i != len(words)-1 {
				for i := 0; i < eachSpaceNum; i++ {
					result.WriteByte(' ')
				}
			}
		}
		for i := 0; i < spaceNum%(wordNum-1); i++ {
			result.WriteByte(' ')
		}
	}
	return result.String()
}
