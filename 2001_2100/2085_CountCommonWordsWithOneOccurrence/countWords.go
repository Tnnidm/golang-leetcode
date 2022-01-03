package countWords

func countWords(words1 []string, words2 []string) int {
	wordMap := make(map[string]int)
	for i := 0; i < len(words1); i++ {
		wordMap[words1[i]] = wordMap[words1[i]] + 10000
	}
	for i := 0; i < len(words2); i++ {
		wordMap[words2[i]]++
	}
	result := 0
	for _, v := range wordMap {
		if v == 10001 {
			result++
		}
	}
	return result
}
