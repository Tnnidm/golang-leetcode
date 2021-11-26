package uniqueMorseRepresentations

import "strings"

func uniqueMorseRepresentations(words []string) int {
	morseCounter := make(map[string]int)
	morseMap := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for _, word := range words {
		morseCode := strings.Builder{}
		for _, char := range word {
			morseCode.WriteString(morseMap[char-'a'])
		}
		morseCounter[morseCode.String()]++
	}
	return len(morseCounter)
}
