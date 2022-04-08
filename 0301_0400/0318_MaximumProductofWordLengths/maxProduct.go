package maxProduct

func maxProduct(words []string) int {
	wordNum := len(words)
	charInt := make([]int, wordNum)
	wordLen := make([]int, wordNum)
	for i := 0; i < wordNum; i++ {
		wordLen[i] = len(words[i])
		for j := 0; j < wordLen[i]; j++ {
			charInt[i] |= 1 << (words[i][j] - 'a')
		}
	}
	product := 0
	for i := 0; i < wordNum; i++ {
		for j := i + 1; j < wordNum; j++ {
			if charInt[i]&charInt[j] == 0 && wordLen[i]*wordLen[j] > product {
				product = wordLen[i] * wordLen[j]
			}
		}
	}
	return product
}
