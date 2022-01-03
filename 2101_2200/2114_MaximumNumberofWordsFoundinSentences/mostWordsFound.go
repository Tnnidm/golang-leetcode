package mostWordsFound

func mostWordsFound(sentences []string) int {
	maxWordNum := 0
	for i := 0; i < len(sentences); i++ {
		wordNum := 0
		for _, char := range sentences[i] {
			if char == ' ' {
				wordNum++
			}
		}
		if maxWordNum < wordNum {
			maxWordNum = wordNum
		}
	}
	return maxWordNum + 1
}
