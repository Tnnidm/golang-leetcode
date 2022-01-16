package isSumEqual

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return string2Num(firstWord)+string2Num(secondWord) == string2Num(targetWord)
}

func string2Num(word string) int {
	result := 0
	for _, v := range word {
		result = result*10 + int(v-'a')
	}
	return result
}
