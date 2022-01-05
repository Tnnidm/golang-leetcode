package countValidWords

import "strings"

func countValidWords(sentence string) int {
	tokens := strings.Fields(sentence)
	result := 0
	for _, token := range tokens {
		valid := true
		tokenLen := len(token)
		// condition 1
		for _, char := range token {
			if !((char >= 'a' && char <= 'z') || char == '-' || char == '!' || char == '.' || char == ',') {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		// condition 2
		lineNum := 0
		for i, char := range token {
			if char == '-' {
				lineNum++
				if lineNum > 1 || !strings.ContainsAny(token[i+1:], "abcdefghijklmnopqrstuvwxyz") || !strings.ContainsAny(token[:i], "abcdefghijklmnopqrstuvwxyz") {
					valid = false
					break
				}
			}
		}
		if !valid {
			continue
		}
		// condition 3
		if strings.Contains(token[:tokenLen-1], "!") || strings.Contains(token[:tokenLen-1], ",") || strings.Contains(token[:tokenLen-1], ".") {
			valid = false
			continue
		}
		result++
	}
	return result
}
