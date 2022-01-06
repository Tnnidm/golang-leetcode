package maxRepeating

import "strings"

func maxRepeating(sequence string, word string) int {
	result := 0
	for repeat := word; ; repeat = repeat + word {
		if !strings.Contains(sequence, repeat) {
			return result
		} else {
			result++
		}
	}
}
