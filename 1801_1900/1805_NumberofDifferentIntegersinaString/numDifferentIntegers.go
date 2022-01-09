package numDifferentIntegers

import (
	"strings"
	"unicode"
)

func numDifferentIntegers(word string) int {
	numStrings := strings.FieldsFunc(word, func(x rune) bool {
		return unicode.IsLetter(x)
	})
	for i := range numStrings {
		for numStrings[i][0] == '0' && len(numStrings[i]) > 1 {
			numStrings[i] = numStrings[i][1:]
		}
	}
	numMap := map[string]struct{}{}
	for _, numString := range numStrings {
		numMap[numString] = struct{}{}
	}
	return len(numMap)
}
